#!/usr/bin/python
#
# Copyright 2019 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8
from __future__ import absolute_import, division, print_function

from marshmallow import fields, validate

from polyaxon.schemas.polyflow.run.replica import (
    ReplicaContainerConfig,
    ReplicaContainerSchema,
)


class ContainerSchema(ReplicaContainerSchema):
    kind = fields.Str(allow_none=True, validate=validate.Equal("container"))

    @staticmethod
    def schema_config():
        return ContainerConfig


class ContainerConfig(ReplicaContainerConfig):
    SCHEMA = ContainerSchema
    IDENTIFIER = "container"

    def __init__(
        self,
        image=None,
        image_pull_policy=None,
        command=None,
        args=None,
        kind=IDENTIFIER,
    ):
        super(ContainerConfig, self).__init__(
            image=image, image_pull_policy=image_pull_policy, command=command, args=args
        )
        self.kind = kind