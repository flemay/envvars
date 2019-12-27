#!/usr/bin/env bash
set -e
set -u

PROFILE=profile.out
[ -f "${PROFILE}" ] || { echo "${PROFILE} not found"; exit 1 ;}
bash <(curl -s https://codecov.io/bash) -f ${PROFILE}