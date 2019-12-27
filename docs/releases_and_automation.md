# Automation & Configuration

## Travis CI

[![Build Status][linkTravisCIProjectBadge]][linkTravisCIProject]

[Travis CI][linkTravisCIProject] is used to test, build and push docker image and behaves a bit differently depending on what triggers the pipeline.

### master branch

master is always releasable (unless tests are not passing). Any changes to master triggers a Travis CI build which test the code, send the code coverage, build an image, but does not push any image.

### Git tag

Any push of a tag will trigger Travis CI build which then tests, builds, and pushes a new Docker image to Docker Hub.

1. Update version in Makefile
1. Update version in README.md
1. Test everything `$ make`
1. Commit the changes and push
1. Run `$ make tag`
1. Go to [Travis CI][linkTravisCIProject] and you should see the build trigger
1. Once the build passed go to [flemay/envvars][linkDockerHubProject] on Docker Hub
1. In `Build Details` tab, you should now see the new Docker image version

### Pull Request

A pull request will trigger Travis CI build which tests the code and the Docker image. It does not push any image.

## Docker Hub

[![Docker Hub][linkDockerHubProjectBadge]][linkDockerHubProject]

Docker Hub is used to store the `flemay/envars` images. The Docker Hub autobuild/autotest are not used for this project because Travis CI gives more control on how the project is tested and built. This process is repeatable/portable with different CI/CD tools (like GitLab) and also for other docker registries.

> Docker multi-stage build could be used to test and build the application but then adding code coverage to the stage just does not feel right. Moreover, I find the code cleaner without the multi-stage. Also, Docker Hub changed his pipeline process and it seems not possible to trigger all the docker builds at once. Lastly, you will know straight away if the image has been pushed successfully from the Travis CI pipeline, so you don't need to look at the docker hub pipeline as well.

The following is a step-by-step guide on how I configured Docker Hub `flemay/envvars`.

1. It is handy to have 2 Docker Hub users: 1 for the creation of the docker registry, the other to push the images
1. Go to [https://hub.docker.com][linkDockerHub] and sign in with your main docker hub user.
1. Go to `Create` and  `Create Automated Build`
1. Select Github
1. Select User `flemay` and then the repository `envvars`
1. Fill out the form (Namespace, Name, Visibility, and Short Description)
1. Go to `Builds`, `Configure Automated Builds`, delete all `BUILD RULES` and `Build Triggers`, and save your modification.
1. Go to `Collaborators` tab and add the docker hub user you want to use for automated builds. You can now use this user in Travis CI to login to docker hub and push the image.

## Codecov

[![codecov][linkCodecovProjectBadge]][linkCodecovProject]

The code coverage is uploaded to [Codecov](https://travis-ci.org/flemay/envvars) after a successful Travis CI build.

## Go Report Card

[![Go Report Card][linkGoReportCardProjectBadge]][linkGoReportCardProject]

[Go Report Card][linkGoReportCardProject] reports the quality of envvars


[linkTravisCIProjectBadge]: https://travis-ci.org/flemay/envvars.svg?branch=master
[linkTravisCIProject]: https://travis-ci.org/flemay/envvars
[linkDockerHubProjectBadge]: https://img.shields.io/badge/dockerhub-builds-blue.svg
[linkDockerHubProject]: https://hub.docker.com/r/flemay/envvars
[linkDockerHub]: https://hub.docker.com
[linkCodecovProjectBadge]: https://codecov.io/gh/flemay/envvars/branch/master/graph/badge.svg
[linkCodecovProject]: https://codecov.io/gh/flemay/envvars
[linkGoReportCardProjectBadge]: https://goreportcard.com/badge/github.com/flemay/envvars
[linkGoReportCardProject]: https://goreportcard.com/report/github.com/flemay/envvars