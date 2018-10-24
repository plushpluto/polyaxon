import uuid

from django.conf import settings
from django.contrib.postgres.fields import JSONField
from django.db import models
from django.utils import timezone
from django.utils.functional import cached_property

import auditor

from constants.experiments import ExperimentLifeCycle
from constants.jobs import JobLifeCycle
from db.models.abstract_jobs import TensorboardJobMixin
from db.models.charts import ChartViewModel
from db.models.cloning_strategies import CloningStrategy
from db.models.statuses import LastStatusMixin, StatusModel
from db.models.unique_names import EXPERIMENT_UNIQUE_NAME_FORMAT
from db.models.utils import (
    DataReference,
    DescribableModel,
    DiffModel,
    NameableModel,
    OutputsModel,
    PersistenceModel,
    ReadmeModel,
    RunTimeModel,
    TagModel
)
from db.redis.heartbeat import RedisHeartBeat
from event_manager.events.experiment import (
    EXPERIMENT_COPIED,
    EXPERIMENT_RESTARTED,
    EXPERIMENT_RESUMED
)
from libs.spec_validation import validate_experiment_spec_config
from schemas.specifications import ExperimentSpecification
from schemas.tasks import TaskType


def default_run_env():
    return {'in_cluster': True}


class Experiment(DiffModel,
                 RunTimeModel,
                 NameableModel,
                 DataReference,
                 OutputsModel,
                 PersistenceModel,
                 DescribableModel,
                 ReadmeModel,
                 TagModel,
                 LastStatusMixin,
                 TensorboardJobMixin):
    """A model that represents experiments."""
    STATUSES = ExperimentLifeCycle

    uuid = models.UUIDField(
        default=uuid.uuid4,
        editable=False,
        unique=True,
        null=False)
    project = models.ForeignKey(
        'db.Project',
        on_delete=models.CASCADE,
        related_name='experiments')
    user = models.ForeignKey(
        settings.AUTH_USER_MODEL,
        on_delete=models.CASCADE,
        related_name='+')
    build_job = models.ForeignKey(
        'db.BuildJob',
        on_delete=models.SET_NULL,
        blank=True,
        null=True,
        related_name='experiments')
    experiment_group = models.ForeignKey(
        'db.ExperimentGroup',
        on_delete=models.CASCADE,
        blank=True,
        null=True,
        related_name='experiments',
        help_text='The experiment group that generate this experiment.')
    declarations = JSONField(
        blank=True,
        null=True,
        help_text='The parameters used for this experiment.')
    config = JSONField(
        null=True,
        blank=True,
        help_text='The compiled polyaxon with specific values for this experiment.',
        validators=[validate_experiment_spec_config])
    run_env = JSONField(
        blank=True,
        null=True,
        default=default_run_env,
        help_text='The run environment of the experiment.')
    original_experiment = models.ForeignKey(
        'self',
        on_delete=models.SET_NULL,
        null=True,
        blank=True,
        related_name='clones',
        help_text='The original experiment that was cloned from.')
    cloning_strategy = models.CharField(
        max_length=16,
        blank=True,
        null=True,
        choices=CloningStrategy.CHOICES)
    status = models.OneToOneField(
        'db.ExperimentStatus',
        related_name='+',
        blank=True,
        null=True,
        editable=True,
        on_delete=models.SET_NULL)
    last_metric = JSONField(
        blank=True,
        null=True,
        default=dict)
    code_reference = models.ForeignKey(
        'db.CodeReference',
        on_delete=models.SET_NULL,
        blank=True,
        null=True,
        related_name='+')

    class Meta:
        app_label = 'db'
        unique_together = (('project', 'name'),)

    @property
    def unique_name(self):
        if self.experiment_group:
            parent_name = self.experiment_group.unique_name
        else:
            parent_name = self.project.unique_name
        return EXPERIMENT_UNIQUE_NAME_FORMAT.format(parent_name=parent_name, id=self.id)

    @cached_property
    def specification(self):
        return ExperimentSpecification(values=self.config) if self.config else None

    @cached_property
    def secret_refs(self):
        return self.specification.secret_refs

    @cached_property
    def configmap_refs(self):
        return self.specification.configmap_refs

    @cached_property
    def resources(self):
        if not self.specification:
            return None
        return self.specification.total_resources

    @cached_property
    def framework(self):
        return self.specification.framework

    @property
    def last_job_statuses(self):
        """The last constants of the job in this experiment."""
        statuses = []
        for status in self.jobs.values_list('status__status', flat=True):
            if status is not None:
                statuses.append(status)
        return statuses

    @property
    def has_running_jobs(self):
        """"Return a boolean indicating if the experiment has any running jobs"""
        return self.jobs.exclude(status__status__in=ExperimentLifeCycle.DONE_STATUS).exists()

    @property
    def calculated_status(self):
        master_status = self.jobs.filter(role=TaskType.MASTER)[0].last_status
        calculated_status = master_status if JobLifeCycle.is_done(master_status) else None
        if calculated_status is None:
            calculated_status = ExperimentLifeCycle.jobs_status(self.last_job_statuses)
        if calculated_status is None:
            return self.last_status
        return calculated_status

    @property
    def is_clone(self):
        return self.original_experiment is not None

    @property
    def original_unique_name(self):
        return self.original_experiment.unique_name if self.original_experiment else None

    @property
    def is_restart(self):
        return self.is_clone and self.cloning_strategy == CloningStrategy.RESTART

    @property
    def is_resume(self):
        return self.is_clone and self.cloning_strategy == CloningStrategy.RESUME

    @property
    def is_copy(self):
        return self.is_clone and self.cloning_strategy == CloningStrategy.COPY

    @property
    def is_independent(self):
        """If the experiment belongs to a experiment_group or is independently created."""
        return self.experiment_group is None

    def update_status(self):
        current_status = self.last_status
        calculated_status = self.calculated_status
        if calculated_status != current_status:
            # Add new status to the experiment
            self.set_status(calculated_status)
            return True
        return False

    def set_status(self, status, message=None, traceback=None, **kwargs):
        if status in ExperimentLifeCycle.HEARTBEAT_STATUS:
            RedisHeartBeat.experiment_ping(self.id)
        if ExperimentLifeCycle.can_transition(status_from=self.last_status, status_to=status):
            ExperimentStatus.objects.create(experiment=self,
                                            status=status,
                                            message=message,
                                            traceback=traceback)

    def _clone(self,
               cloning_strategy,
               event_type,
               user=None,
               description=None,
               config=None,
               declarations=None,
               code_reference=None,
               update_code_reference=False,
               experiment_group=None):
        if not code_reference and not update_code_reference:
            code_reference = self.code_reference
        instance = Experiment.objects.create(
            project=self.project,
            user=user or self.user,
            experiment_group=experiment_group,
            description=description or self.description,
            config=config or self.config,
            declarations=declarations or self.declarations,
            original_experiment=self,
            cloning_strategy=cloning_strategy,
            code_reference=code_reference)
        auditor.record(event_type=event_type, instance=instance)
        return instance

    def resume(self,
               user=None,
               description=None,
               config=None,
               declarations=None,
               code_reference=None,
               update_code_reference=False,
               experiment_group=None):
        # TODO: We need to check if this instance was stopped after it was created
        # TODO: If that's the case and no updates are passed we just resume this very same instance
        # If the current instance is a resume of an original than we need to resume the original
        if self.is_resume:
            return self.original_experiment.resume(
                user=user,
                description=description,
                config=config,
                declarations=declarations,
                code_reference=code_reference,
                update_code_reference=update_code_reference,
                experiment_group=experiment_group or self.experiment_group)

        # Resume normal workflow
        return self._clone(cloning_strategy=CloningStrategy.RESUME,
                           event_type=EXPERIMENT_RESUMED,
                           user=user,
                           description=description,
                           config=config,
                           declarations=declarations,
                           code_reference=code_reference,
                           update_code_reference=update_code_reference,
                           experiment_group=experiment_group or self.experiment_group)

    def restart(self,
                user=None,
                description=None,
                config=None,
                declarations=None,
                code_reference=None,
                update_code_reference=False,
                experiment_group=None):
        return self._clone(cloning_strategy=CloningStrategy.RESTART,
                           event_type=EXPERIMENT_RESTARTED,
                           user=user,
                           description=description,
                           config=config,
                           declarations=declarations,
                           code_reference=code_reference,
                           update_code_reference=update_code_reference,
                           experiment_group=experiment_group)

    def copy(self,
             user=None,
             description=None,
             config=None,
             declarations=None,
             code_reference=None,
             update_code_reference=False,
             experiment_group=None):
        return self._clone(cloning_strategy=CloningStrategy.COPY,
                           event_type=EXPERIMENT_COPIED,
                           user=user,
                           description=description,
                           config=config,
                           declarations=declarations,
                           code_reference=code_reference,
                           update_code_reference=update_code_reference,
                           experiment_group=experiment_group)


class ExperimentStatus(StatusModel):
    """A model that represents an experiment status at certain time."""
    STATUSES = ExperimentLifeCycle

    experiment = models.ForeignKey(
        'db.Experiment',
        on_delete=models.CASCADE,
        related_name='statuses')
    status = models.CharField(
        max_length=64,
        blank=True,
        null=True,
        default=STATUSES.CREATED,
        choices=STATUSES.CHOICES)

    class Meta:
        app_label = 'db'
        verbose_name_plural = 'Experiment Statuses'
        ordering = ['created_at']

    def __str__(self):
        return '{} <{}>'.format(self.experiment.unique_name, self.status)


class ExperimentMetric(models.Model):
    """A model that represents an experiment metric at certain time."""
    experiment = models.ForeignKey(
        'db.Experiment',
        on_delete=models.CASCADE,
        related_name='metrics')
    created_at = models.DateTimeField(default=timezone.now)
    values = JSONField()

    def __str__(self):
        return '{} <{}>'.format(self.experiment.unique_name, self.created_at)

    class Meta:
        app_label = 'db'
        ordering = ['created_at']


class ExperimentChartView(ChartViewModel):
    """A model that represents an experiment chart view."""
    experiment = models.ForeignKey(
        'db.Experiment',
        on_delete=models.CASCADE,
        related_name='chart_views')

    class Meta:
        app_label = 'db'
        verbose_name_plural = 'Experiment Chart Views'
        ordering = ['created_at']
