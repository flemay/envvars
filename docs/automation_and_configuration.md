# Automation & Configuration

## Travis CI

[![Build Status](https://travis-ci.org/flemay/envvars.svg?branch=master)](https://travis-ci.org/flemay/envvars)

[Travis CI](https://travis-ci.org/flemay/envvars) is used to test envvars whenever a code change is committed to master. It uses the [3 Musketeer](https://github.com/flemay/three-musketeers).

## Docker Hub

[![Docker Build Status](https://img.shields.io/docker/build/flemay/envvars.svg)](https://hub.docker.com/r/flemay/envvars)

The following is a step-by-step guide on how I configured Docker Hub `flemay/envvars`.

1. Go to [https://hub.docker.com](https://hub.docker.com) and sign in
1. Go to `Create` and  `Create Automated Build`
1. Select Github
1. Select User `flemay` and then the repository `envvars`
1. Fill out the form (Namespace, Name, Visibility, and Short Description)
1. Since the builds will be based on tags only, click on `Click here to customize`
1. There are 2 Autobuild Tags
    1. For tag
        1. Push Type: `Tag`
        1. Name: `/.*/`
        1. Dockerfile Location: `/`
        1. Docker Tag: _empty_
    1. For Branch master
        1. Push Type: `Branch`
        1. Name: `master`
        1. Dockerfile Location: `/`
        1. Docker Tag: `latest`
1. Uncheck "When active, builds will happen automatically on pushes"
1. Click on button `Create`

## Codecov

[![codecov](https://codecov.io/gh/flemay/envvars/branch/master/graph/badge.svg)](https://codecov.io/gh/flemay/envvars)

The code coverage is uploaded to [Codecov](https://travis-ci.org/flemay/envvars) after a successful Travis CI build.

## Go Report Card

[![Go Report Card](https://goreportcard.com/badge/github.com/flemay/envvars)](https://goreportcard.com/report/github.com/flemay/envvars)

[Go Report Card](https://goreportcard.com/report/github.com/flemay/envvars) reports the quality of envvars
