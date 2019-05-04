#!/usr/bin/env sh
set -e
set -u

# For more Travis CI environment variables https://docs.travis-ci.com/user/environment-variables/

if [ ! -z "${TRAVIS_TAG}" ]; then
  echo "Triggered on Git Tag"
  if [ "${GIT_TAG}" == "${TRAVIS_TAG}" ]; then
    make onGitTag
  else
    echo "Error: TRAVIS_TAG '${TRAVIS_TAG}' differs from GIT_TAG '${GIT_TAG}'"
    exit 1
  fi
elif [ ! -z "${TRAVIS_PULL_REQUEST}" ]; then
  echo "Triggered on Pull Request"
  make onPullRequest
elif [ "${TRAVIS_BRANCH}" = "master" ]; then
  echo "Triggered on Commit/Merge/Schedule to Master"
  make onMasterChange
else
  echo "Error: This case is not handled!"
  exit 1
fi