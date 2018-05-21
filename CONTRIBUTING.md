# Contributing to Envvars

Envvars is an open source project and contributions are greatly appreciated.

There are few ways to contribute to this project.

## Star it

If you like the project, go on [GitHub](https://github.com/flemay/envvars) and ⭐️ it!

## Share it

If you think Envvars can benefit your friends, teammates, and company, share it!

## Feedback

Feedback is greatly appreciated. Do you have workflows that the tool supports well, or doesn't support at all? Do any of the commands have surprising effects, output, or results? Let us know by filing an issue, describing what you did or wanted to do, what you expected to happen, and what actually happened.

# Contributing code

The project follows the typical GitHub pull request model. Before starting any work, please either comment on an existing issue, or file a new one.

## 1. Fork

[Fork Envvars](https://github.com/flemay/envvars/fork) on GitHub so that you can commit your changes and create a Pull Request.

## 2. Development

There are 2 approaches (described below) to make your changes:

- [3 Musketeers](https://github.com/flemay/3musketeers) which uses Make, Docker, and Compose
- a Go environment with dep

## 3. Create a Pull Request

Happy with all your precious changes? Create a [Pull Request](https://help.github.com/articles/creating-a-pull-request/)!

# Development with the 3 Musketeers

## Prerequisites

- Docker
- Compose
- Make

## Steps

```bash
# clone the fork locally (location does not matter)
$ git clone git@github.com:username/envvars.git
$ cd envvars
# checkout a new branch
$ git checkout -b meaningful_branch_name

# download all the dependencies
$ make deps
# make your changes
# ...
# test your changes
$ make test
# build Envvars
$ make build
# run Envvars
$ make run

# push your changes
$ git push origin meaningful_branch_name
```

# Development with Go environment

## Prerequisites

- Go 1.10
- [dep](https://github.com/golang/dep)
- Make

## Steps

```bash
# create folder
$ mkdir -p path/to/go/src/github.com/flemay
$ cd path/to/go/src/github.com/flemay
# clone the fork locally
$ git clone git@github.com:username/envvars.git
$ cd envvars
# checkout a new branch
$ git checkout -b meaningful_branch_name

# download all the dependencies
$ make _deps
# make your changes
# ...
# test your changes
$ make _test
# build Envvars
$ make _build
# run Envvars
$ make _run

# push your changes
$ git push origin meaningful_branch_name
```

# Mocks

[Mockery](https://github.com/vektra/mockery) is being used to generate mocks in the folder `pkg/mocks`.

- Mockery must be installed locally
- To generate all of them `$ make _mock`
- Once generated, one would need to fix all the mocks imports

# Travis CI & Docker Hub

Travis CI is used to test, build and trigger Docker Hub to build the Docker Image.

> This assumes master contains the latest changes to be released.

## Master and Latest

Any changes to master branch will triggers Travis CI build which then call Docker Hub to build the `latest` image.

1. Go to [Travis CI](https://travis-ci.org/flemay/envvars) and you should see the build trigger
1. Once the build passed go to [flemay/envvars](https://hub.docker.com/r/flemay/envvars) on Docker Hub
1. In `Build Details` tab, you should now see the build kicking off

## Tag release

Any push of a tag will trigger Travis CI build which then call Docker Hub to build the tag image

1. Update version in Makefile
1. Update version in README.md
1. Build the image locally `$ make dockerBuild`
1. Test the image locally `$ make dockerTest`
1. Commit the changes and push
1. Run `$ make tag`
1. Go to [Travis CI](https://travis-ci.org/flemay/envvars) and you should see the build trigger
1. Once the build passed go to [flemay/envvars](https://hub.docker.com/r/flemay/envvars) on Docker Hub
1. In `Build Details` tab, you should now see the build kicking off

## Monthly update

There is a cron task in Travis CI to run the build which will trigger Docker Hub to rebuild all the images (latest and tags).
