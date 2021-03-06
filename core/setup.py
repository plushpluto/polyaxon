#!/usr/bin/env python
import os
import sys

from setuptools import find_packages, setup
from setuptools.command.test import test as TestCommand


def read_readme():
    if not os.path.exists("../README.md"):
        return ""
    with open("../README.md") as f:
        return f.read()


class PyTest(TestCommand):
    def finalize_options(self):
        TestCommand.finalize_options(self)
        self.test_args = []
        self.test_suite = True

    def run_tests(self):
        import pytest

        errcode = pytest.main(self.test_args)
        sys.exit(errcode)


setup(
    name="polyaxon",
    version="1.0.87",
    description="Command Line Interface (CLI) and client to interact with Polyaxon API.",
    long_description=read_readme(),
    long_description_content_type="text/markdown",
    maintainer="Polyaxon, Inc.",
    maintainer_email="contact@polyaxon.com",
    author="Polyaxon, Inc.",
    author_email="contact@polyaxon.com",
    url="https://github.com/polyaxon/polyaxon",
    license="Apache 2.0",
    platforms="any",
    packages=find_packages(),
    keywords=[
        "polyaxon",
        "aws",
        "s3",
        "microsoft",
        "azure",
        "google cloud storage",
        "gcs",
        "deep-learning",
        "machine-learning",
        "data-science",
        "neural-networks",
        "artificial-intelligence",
        "ai",
        "reinforcement-learning",
        "kubernetes",
        "aws",
        "microsoft",
        "azure",
        "google cloud",
        "tensorFlow",
        "pytorch",
    ],
    install_requires=[
        "click<8.0",
        "click-completion<0.6",
        "tabulate>=0.8.2",
        "Jinja2>=2.10.3",
        "kubernetes>=10.0.1",
        "marshmallow==3.4.0",
        "polyaxon-sdk==1.0.87",
        "python-dateutil>=2.7.3",
        "pytz>=2019.2",
        "PyYAML>=5.1",
        "ujson>=1.35",
        "requests>=2.20.1",
        "requests-toolbelt>=0.8.0",
        "sentry-sdk>=0.12.3",
        "urllib3>=1.25.6",
        "certifi>=2019.9.11",
        "psutil",
        "nvidia-ml-py3",
    ],
    extras_require={
        "gcs": ["google-cloud-storage"],
        "s3": ["boto3", "botocore"],
        "azure": ["azure-storage-blob>=12.3.1"],
        "docker": ["docker"],
        "git": ["git"],
        "polytune": ["scikit-learn==0.22.2", "hyperopt==0.2.4"],
        "polyboard": ["Pillow", "matplotlib", "moviepy", "plotly", "bokeh", "pandas"],
        "streams": [
            "kubernetes-asyncio==11.2.0",
            "starlette==0.13.4",
            "aiofiles==0.5.0",
            "uvicorn==0.11.5",
            "pandas",
        ],
    },
    entry_points={"console_scripts": ["polyaxon = polyaxon.main:cli"]},
    python_requires=">=3.5",
    classifiers=[
        "Programming Language :: Python",
        "Programming Language :: Python :: 3",
        "Programming Language :: Python :: 3.5",
        "Programming Language :: Python :: 3.6",
        "Programming Language :: Python :: 3.7",
        "Programming Language :: Python :: 3.8",
        "Operating System :: OS Independent",
        "Intended Audience :: Developers",
        "Intended Audience :: Science/Research",
        "Topic :: Scientific/Engineering :: Artificial Intelligence",
    ],
    tests_require=["pytest"],
    cmdclass={"test": PyTest},
)
