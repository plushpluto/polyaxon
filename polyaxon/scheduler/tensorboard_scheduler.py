import logging

from django.conf import settings
from kubernetes.client.rest import ApiException

from constants.jobs import JobLifeCycle
from scheduler.spawners.tensorboard_spawner import TensorboardSpawner
from scheduler.spawners.utils import get_job_definition

logger = logging.getLogger('polyaxon.scheduler.tensorboard')


def start_tensorboard(project):
    spawner = TensorboardSpawner(
        project_name=project.unique_name,
        project_uuid=project.uuid.hex,
        k8s_config=settings.K8S_CONFIG,
        namespace=settings.K8S_NAMESPACE,
        in_cluster=True)

    try:
        results = spawner.start_tensorboard(image=project.tensorboard.image,
                                            resources=project.tensorboard.resources,
                                            node_selectors=project.tensorboard.node_selectors)
    except ApiException as e:
        logger.warning('Could not start tensorboard, please check your polyaxon spec %s', e)
        project.tensorboard.set_status(
            JobLifeCycle.FAILED,
            message='Could not start tensorboard, encountered a Kubernetes ApiException.')
        return
    except Exception as e:
        logger.warning('Could not start tensorboard, please check your polyaxon spec %s', e)
        project.tensorboard.set_status(
            JobLifeCycle.FAILED,
            message='Could not start tensorboard encountered an {} exception.'.format(
                e.__class__.__name__
            ))
        return
    project.tensorboard.definition = get_job_definition(results)
    project.tensorboard.save()


def stop_tensorboard(project, update_status=False):
    spawner = TensorboardSpawner(
        project_name=project.unique_name,
        project_uuid=project.uuid.hex,
        k8s_config=settings.K8S_CONFIG,
        namespace=settings.K8S_NAMESPACE,
        in_cluster=True)

    spawner.stop_tensorboard()
    if update_status:
        # Update experiment status to show that its stopped
        project.tensorboard.set_status(status=JobLifeCycle.STOPPED,
                                       message='Tensorboard was stopped')


def get_tensorboard_url(project):
    spawner = TensorboardSpawner(
        project_name=project.unique_name,
        project_uuid=project.uuid.hex,
        k8s_config=settings.K8S_CONFIG,
        namespace=settings.K8S_NAMESPACE,
        in_cluster=True)
    return spawner.get_tensorboard_url()
