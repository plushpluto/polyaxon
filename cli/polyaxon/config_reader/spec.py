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

import json
import os
import six
import yaml

from collections import Mapping
from yaml.parser import ParserError  # noqa
from yaml.scanner import ScannerError  # noqa

from polyaxon.exceptions import PolyaxonSchemaError


class ConfigSpec(object):
    def __init__(self, value, config_type=None, check_if_exists=True):
        self.value = value
        self.config_type = config_type
        self.check_if_exists = check_if_exists

    @classmethod
    def get_from(cls, value):
        if isinstance(value, cls):
            return value

        return cls(value=value)

    def check_type(self):
        type_check = self.config_type is None and not isinstance(
            self.value, (Mapping, six.string_types)
        )
        if type_check:
            raise PolyaxonSchemaError(
                "Expects Mapping, string, or list of Mapping/string instances, "
                "received {} instead".format(type(self.value))
            )

    def read(self):
        if isinstance(self.value, Mapping):
            config_results = self.value
        elif os.path.isfile(self.value):
            config_results = _read_from_file(self.value, self.config_type)
        else:
            # try reading a stream of yaml or json
            try:
                config_results = _read_from_stream(self.value)
            except (ScannerError, ParserError):
                raise PolyaxonSchemaError(
                    "Received non valid yaml stream: `{}`".format(self.value)
                )
        return config_results


def _read_from_stream(stream):
    results = _read_from_yml(stream, is_stream=True)
    if not results:
        results = _read_from_json(stream, is_stream=True)
    return results


def _read_from_file(f_path, file_type):
    _, ext = os.path.splitext(f_path)
    if ext in (".yml", ".yaml") or file_type in (".yml", ".yaml"):
        return _read_from_yml(f_path)
    elif ext == ".json" or file_type == ".json":
        return _read_from_json(f_path)
    raise PolyaxonSchemaError(
        "Expects a file with extension: `.yml`, `.yaml`, or `json`, "
        "received instead `{}`".format(ext)
    )


def _read_from_yml(f_path, is_stream=False):
    try:
        if is_stream:
            return yaml.safe_load(f_path)
        with open(f_path) as f:
            return yaml.safe_load(f)
    except (ScannerError, ParserError):
        raise PolyaxonSchemaError("Received non valid yaml: `{}`".format(f_path))


def _read_from_json(f_path, is_stream=False):
    if is_stream:
        try:
            return json.loads(f_path)
        except ValueError as e:
            raise PolyaxonSchemaError(e)
    try:
        return json.loads(open(f_path).read())
    except ValueError as e:
        raise PolyaxonSchemaError(e)
