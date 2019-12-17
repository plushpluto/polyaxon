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

import os

from polyaxon.exceptions import PolyaxonStoresException
from polyaxon.stores.stores.base_store import BaseStore


class StoreManager(object):
    """
    A convenient class to map experiment/job outputs/data paths to a given/configured store.
    """

    def __init__(self, store=None, path=None):
        self._path = path
        if not store and path:
            store = BaseStore.get_store_for_path(path=path)
        if not store:
            store = BaseStore.get_store()
        if isinstance(store, BaseStore):
            self._store = store
        else:
            raise PolyaxonStoresException(
                "Received an unrecognised store `{}`.".format(store)
            )

    @classmethod
    def get_for_type(cls, store_type, store_access):
        store = BaseStore.get_store_for_type(
            store_type=store_type, store_access=store_access
        )
        return cls(store=store)

    def set_store(self, store):
        self._store = store

    def set_path(self, path):
        self._path = path

    def set_env_vars(self):
        """Set authentication and access of the current store to the env vars"""
        if self.store:
            self.store.set_env_vars()

    @property
    def store(self):
        return self._store

    @property
    def path(self):
        return self._path

    def ls(self, path, sort=True):
        if self._path:  # We assume rel paths
            path = os.path.join(self._path, path)
        results = self.store.ls(path)
        if sort:
            return {"files": sorted(results["files"]), "dirs": sorted(results["dirs"])}
        return results

    def list(self, path):
        if self._path:  # We assume rel paths
            path = os.path.join(self._path, path)
        return self.store.list(path)

    def delete(self, path):
        return self.store.delete(path)

    def upload_file(self, filename, path=None, **kwargs):
        path = path or self._path
        self.store.upload_file(filename, path, **kwargs)

    def upload_dir(self, dirname, path=None, **kwargs):
        path = path or self._path
        self.store.upload_dir(dirname, path, **kwargs)

    def download_file(self, filename, local_path=None, use_basename=False, **kwargs):
        if self._path:  # We assume rel paths
            file_path = os.path.join(self._path, filename)
            local_path = local_path or filename
        else:
            file_path = filename
        self.store.download_file(
            file_path, local_path, use_basename=use_basename, **kwargs
        )

    def download_dir(self, dirname, local_path=None, use_basename=False, **kwargs):
        if self._path:  # We assume rel paths
            dir_path = os.path.join(self._path, dirname)
            local_path = local_path or dirname
        else:
            dir_path = dirname
        self.store.download_dir(
            dir_path, local_path, use_basename=use_basename, **kwargs
        )
