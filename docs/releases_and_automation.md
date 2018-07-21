# Automation & Configuration

## Releases

Master is always releasable (unless tests are not passing).

### Tag

Any push of a tag will trigger Travis CI build which then call Docker Hub to build the tag image

1. Update version in Makefile
1. Update version in README.md
1. Build the image locally `$ make dockerBuild`
1. Test the image locally `$ make dockerTest`
1. Commit the changes and push
1. Run `$ make _tag`
1. Go to [Travis CI](https://travis-ci.org/flemay/envvars) and you should see the build trigger
1. Once the build passed go to [flemay/envvars](https://hub.docker.com/r/flemay/envvars) on Docker Hub
1. In `Build Details` tab, you should now see the build kicking off

### Monthly update

There is a cron task in Travis CI to run the build which will trigger Docker Hub to rebuild all the images (latest and tags).

## Travis CI

[![Build Status](https://travis-ci.org/flemay/envvars.svg?branch=master)](https://travis-ci.org/flemay/envvars)

[Travis CI](https://travis-ci.org/flemay/envvars) tests envvars whenever a code change is committed to master. It uses the [3 Musketeer](https://github.com/flemay/3musketeers).

Once the test passed, Travis CI triggers a Docker Hub build for building a new image with the tag latest.

1. Go to [Travis CI](https://travis-ci.org/flemay/envvars) and you should see the build trigger
1. Once the build passed go to [flemay/envvars](https://hub.docker.com/r/flemay/envvars) on Docker Hub
1. In `Build Details` tab, you should now see the build kicking off

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
