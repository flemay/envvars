#!/usr/bin/env bash

set -e
set -u

echo "GITHUB_SHA: ${GITHUB_SHA}"
echo "GITHUB_REF: ${GITHUB_REF}"
echo "TAG: ${TAG}"

[[ "${GITHUB_REF}" == "refs/tags/"* ]] || { echo "Not from a tag created event. Skip release."; exit 0; }

GITHUB_TAG=$(echo ${GITHUB_REF} | cut -d'/' -f 3)
echo "GITHUB_TAG: ${GITHUB_TAG}"

[ "${TAG}" == "${GITHUB_TAG}" ] || { echo "TAG differs from GITHUB_TAG"; exit 1; }

make envfile deps test build run buildDockerImage pushDockerImage clean
