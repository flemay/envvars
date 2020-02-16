# Automation & Configuration

## GitHub

## master branch

master is always releasable (unless tests are not passing). A change to master triggers a GitHub Action which tests the code, sends the code coverage, and builds a Docker image without pushing it.

The master branch is also protected by:

- requiring status checks to pass before merging
  - codecov/project
  - Test GitHub Action
- requiring branches to be up to date before merging

See [Configuring GitHub protected branches][linkConfiguringGitHubProtectedBranches].

## GitHub Actions

[![Build Status][linkGitHubActionsProjectTestBadge]][linkGitHubActionsProject]
[![Build Status][linkGitHubActionsProjectReleaseBadge]][linkGitHubActionsProject]

This project uses [GitHub Actions][linkGitHubActionsProject] to test, build, and push envvars Docker image. There are two: Test and Release.

`Test` is triggered whenever there is a [Pull Request created][linkGitHubActionsPullRequestEvent] or a [change to master branch][linkGitHubActionsPushEvent].

`Release` gets triggered under a [GitHub create event][linkGitHubActionsCreateEvent]. However, it is only on tag created that the release of the Docker image happens.

> [GitHub release event][linkGitHubActionsReleaseEvent] could eventually be used but for now, the release is done on Git tag.

Environment variables nammed `DOCKER_PASSWORD` and `CODECOV_TOKEN` were set in the Secrets section of the repository's settings.

## Git tag

[![GitHub Tag][linkGitHubProjectTagBadge]][linkGitHubProject]

A push of a tag triggers a GitHub Action which tests, builds, and pushes a new Docker image to Docker Hub.

1. Update version in Makefile
1. Update version in README.md
1. Test everything `$ make`
1. Commit the changes and push
1. Run `$ make tag`
1. Go to [GitHub Actions][linkGitHubActionsProject] and you should see the build triggered
1. Once the build passed, go to [flemay/envvars][linkDockerHubProject] on Docker Hub
1. In `Build Details` tab, you should see the new Docker image version

## Docker Hub

[![Docker Hub][linkDockerHubProjectBadge]][linkDockerHubProject]
[![Docker Hub Pulls Badge][LinkDockerHubProjectPullsBadge]][linkDockerHubProject]

Docker Hub is used to store `flemay/envars` images. The Docker Hub autobuild/autotest are not used for this project because GitHub Actions gives more control on how the project is tested and built. This process is repeatable/portable with different CI/CD tools (like GitLab) and also for other docker registries.

> Docker multi-stage build could be used to test and build the application but then adding code coverage to the stage just does not feel right. Moreover, I find the code cleaner without the multi-stage. Also, Docker Hub changed his pipeline process and it seems not possible to trigger all the docker builds at once. Lastly, you will know straight away if the image has been pushed successfully from the GitHub Actions, so you don't need to look at the docker hub pipeline as well.

The following is a step-by-step guide on how I configured Docker Hub `flemay/envvars`.

1. It is handy to have 2 Docker Hub users: 1 for the creation of the docker registry, the other to push the images
1. Go to [https://hub.docker.com][linkDockerHub] and sign in with your main docker hub user.
1. Go to `Repositories`
1. Click `Create Repository` button
1. Fill out the form (Namespace, Name, Description, Visibility)
1. Leave the `Build Settings` section empty
1. Click `Create` button
1. In the `General` tab, update the `Readme` section with this repository's README
1. In the `Collaborators` tab, add the docker hub user you want to use for automated builds. This user can now be used in GitHub Actions to push the image.

> For DOCKER_PASSWORD, it is recommended to use an Access Token and not the password from the automation user. To create it, log in as the automation Docker user, go to Account Settings. From the Security tab, you will be able to create the token. The name of the token can be something like "GitHub Actions".

## Codecov

[![codecov][linkCodecovProjectBadge]][linkCodecovProject]

The code coverage is uploaded to [Codecov][linkCodecovProject] after a successful GitHub Action build.

## Go Report Card

[![Go Report Card][linkGoReportCardProjectBadge]][linkGoReportCardProject]

[Go Report Card][linkGoReportCardProject] reports the quality of envvars


[linkGitHubActionsProjectTestBadge]: https://github.com/flemay/envvars/workflows/Test/badge.svg
[linkGitHubActionsProjectReleaseBadge]: https://github.com/flemay/envvars/workflows/Release/badge.svg
[linkGitHubActionsProject]: https://github.com/flemay/envvars/actions
[linkDockerHubProjectBadge]: https://img.shields.io/badge/dockerhub-builds-blue.svg
[linkDockerHubProject]: https://hub.docker.com/r/flemay/envvars
[linkDockerHub]: https://hub.docker.com
[linkCodecovProjectBadge]: https://codecov.io/gh/flemay/envvars/branch/master/graph/badge.svg
[linkCodecovProject]: https://codecov.io/gh/flemay/envvars
[linkGoReportCardProjectBadge]: https://goreportcard.com/badge/github.com/flemay/envvars
[linkGoReportCardProject]: https://goreportcard.com/report/github.com/flemay/envvars
[linkGitHubProjectTagBadge]: https://img.shields.io/github/tag/flemay/envvars.svg
[linkGitHubProject]: https://github.com/flemay/envvars
[linkConfiguringGitHubProtectedBranches]: https://help.github.com/en/github/administering-a-repository/configuring-protected-branches
[LinkDockerHubProjectPullsBadge]: https://img.shields.io/docker/pulls/flemay/envvars
[linkGitHubActionsCreateEvent]: https://help.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows#create-event-create
[linkGitHubActionsReleaseEvent]: https://help.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows#release-event-release
[linkGitHubActionsPullRequestEvent]: https://help.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows#pull-request-event-pull_request
[linkGitHubActionsPushEvent]: https://help.github.com/en/actions/automating-your-workflow-with-github-actions/events-that-trigger-workflows#push-event-push
