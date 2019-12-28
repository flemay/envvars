#!/usr/bin/env bash

# I am currently using Create event
# (https://help.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows#release-event-release)
# instead of Release event
# (https://help.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows#release-event-release)
# The Create event can come from Tag or Branch creation and this script makes sure
# the release is only done at tag creation.

set -e
set -u

echo "GITHUB_SHA: ${GITHUB_SHA}"
echo "GITHUB_REF: ${GITHUB_REF}"
echo "TAG: ${TAG}"

[[ "${GITHUB_REF}" == "refs/tags/"* ]] || { echo "Not from a tagging event. Skip release."; exit 0; }

GITHUB_TAG=$(echo ${GITHUB_REF} | cut -d'/' -f 3)
echo "GITHUB_TAG: ${GITHUB_TAG}"

[ "${TAG}" == "${GITHUB_TAG}" ] || { echo "TAG differs from GITHUB_TAG"; exit 1; }

make envfile deps test build run buildDockerImage pushDockerImage clean
