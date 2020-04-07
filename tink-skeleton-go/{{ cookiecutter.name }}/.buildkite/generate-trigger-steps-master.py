#!/usr/bin/env python

from __future__ import print_function
import os

BRANCHES = {
    'oxford-production': {'block': True},
    'oxford-staging': {'block': False},
    'leeds-production': {'block': True},
    'leeds-staging': {'block': True},
    'cornwall-production': {'block': True},
    'cornwall-testing': {'block': True},
    'kirkby-production': {'block': True},
    'kirkby-staging': {'block': True},
    'neston-production': {'block': True},
    'neston-preprod': {'block': True},
    'neston-staging': {'block': True},
    'newport-production': {'block': True},
    'newport-staging': {'block': True},
    'farnham-production': {'block': True},
    'farnham-staging': {'block': True},
    'aggregation-production': {'block': True},
    'aggregation-staging': {'block': True},
}

PROJECTS = {
    '{{ cookiecutter.name }}': {'chart': True, 'salt': False},
}

STEP = """
- name: "Trigger release {branch} {project}"
  trigger: "{pipeline}"
  async: true
  build:
    message: "Release {project} {version} to {branch}"
    commit: "HEAD"
    branch: "{branch}"
    env:
      TINK_PROJECT: "{project}"
      TINK_VERSION: "{version}"
      TINK_BRANCH: "{branch}"
      TINK_BLOCK: "{block}"
      TINK_CHART_REPO: "tink-backend"
      TINK_SALT_DEPLOY: "{salt_deploy}"
      TINK_KUBERNETES_DEPLOY: "{kubernetes_deploy}"
"""

version = os.environ['VERSION']

for project, project_settings in PROJECTS.items():
    for branch, settings in BRANCHES.items():
        if settings.get('block'):
            block = 'true'
        else:
            block = ''

        kubernetes_deploy = project_settings.get('chart', False)
        salt_deploy = project_settings.get('salt', False)

        print(STEP.format(
            pipeline='release-{}'.format(project),
            project=project,
            version=version,
            branch=branch,
            block=block,
            salt_deploy=salt_deploy,
            kubernetes_deploy=kubernetes_deploy,
))
