#!/usr/bin/env sh
set -e
set -u

# For more Travis CI environment variables https://docs.travis-ci.com/user/environment-variables/

echo "TRAVIS_TAG: ${TRAVIS_TAG}"
echo "TRAVIS_PULL_REQUEST: ${TRAVIS_PULL_REQUEST}"
echo "TRAVIS_BRANCH: ${TRAVIS_BRANCH}"

if [ ! -z "${TRAVIS_TAG}" ]; then
  echo "Triggered on Git Tag: ${TRAVIS_TAG}"
  if [ "${GIT_TAG}" == "${TRAVIS_TAG}" ]; then
    make onGitTag
  else
    echo "Error: TRAVIS_TAG '${TRAVIS_TAG}' differs from GIT_TAG '${GIT_TAG}'"
    exit 1
  fi
elif [ "${TRAVIS_PULL_REQUEST}" != "false" ]; then
  echo "Triggered on Pull Request: ${TRAVIS_PULL_REQUEST}"
  make onPullRequest
elif [ "${TRAVIS_BRANCH}" = "master" ]; then
  echo "Triggered on Commit/Merge/Schedule to branch: ${TRAVIS_BRANCH}"
  make onMasterChange
else
  echo "Error: This case is not handled!"
  exit 1
fi