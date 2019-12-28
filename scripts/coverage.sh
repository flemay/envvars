#!/usr/bin/env bash
set -e
set -u

PROFILE=profile.out
[ -f "${PROFILE}" ] || { echo "${PROFILE} not found"; exit 1 ;}
curl --data-binary @codecov.yml https://codecov.io/validate
bash <(curl -s https://codecov.io/bash) -f ${PROFILE}