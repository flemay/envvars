<p align="center"><img src="docs/envvars_gopher.png" width="360"></p>

[![Build Status](https://travis-ci.org/flemay/envvars.svg?branch=master)](https://travis-ci.org/flemay/envvars)
[![Go Report Card](https://goreportcard.com/badge/github.com/flemay/envvars)](https://goreportcard.com/report/github.com/flemay/envvars)
[![codecov](https://codecov.io/gh/flemay/envvars/branch/master/graph/badge.svg)](https://codecov.io/gh/flemay/envvars)
[![Docker Build Status](https://img.shields.io/docker/build/flemay/envvars.svg)](https://hub.docker.com/r/flemay/envvars)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](LICENSE)

# Envvars

Envvars, a command line tool written in Go, provides a way to describe the environment variables of a project and ensures they are defined before testing, building, and deploying. It also generates an env file to be used by other applications such as Docker and Compose.

## 3 Musketeers

Envvars fits nicely with the [3 Musketeers](https://github.com/flemay/3musketeers) for managing the environment variables used by an application.

## Declaration File

The declaration file (written in [YAML](http://yaml.org/spec/1.2/spec.html)) is the core of Envvars. It declares all the environment variables used by a project.

Envvars is looking for the declaration file `envvars.yml` by default. A different file can be passed with the flag `-f path/to/declarationfile.yml`.

```yml
tags:
  - name: deploy
    desc: tag used when deploying

envvars:
  - name: ENV
    desc: Application stage (dev, qa, preprod, prod)
    tags:
      - deploy
    optional: true
    example: dev
```

| Field            |      Type      | Required | Description                                                                                                                        |
|------------------|:--------------:|:--------:|------------------------------------------------------------------------------------------------------------------------------------|
| tags             |      list      |    no    | List of tags to be used for targeting a subset of environment variables                                                            |
| tags.name        |     string     |    yes   | Unique tag name                                                                                                                    |
| tags.desc        |     string     |    yes   | Meaningful description of the tag                                                                                                  |
| envvars          |      list      |    yes   | List of environment variables                                                                                                      |
| envvars.name     |     string     |    yes   | Unique environment variable name                                                                                                   |
| envvars.desc     |     string     |    yes   | Meaningful description of the environment variable                                                                                 |
| envvars.tags     | list of string |    no    | List of tags for the environment variable. Each tag must be declared in the "tags" field.                                          |
| envvars.optional |      bool      |    no    | Allows the environment variable to be empty or not defined. It is best to avoid it unless your application accepts an empty value. |
| envvars.example  |     string     |    no    | Example value for the environment variable.                                                                                        |

## Installation

```bash
# with go get
$ go get -u github.com/flemay/envvars

# or use the tiny docker image (< 5 MB)
$ docker run --rm flemay/envvars
```

## Usage

```bash
# create a declaration file envvars.yml
# envvars:
#   - name: ECHO
#     desc: env var ECHO
$ printf "envvars:\n  - name: ECHO\n    desc: env var ECHO\n" > envvars.yml

# validate the declaration file if it contains errors
$ envvars validate

# ensure the environment variables comply with the declaration file
$ envvars ensure
# Error: environment variable ECHO is not defined
# set ECHO with empty value
$ export ECHO=""
$ envvars ensure
# Error: environment variable ECHO is empty
# set ECHO with non-empty value
$ export ECHO="helloworld"
$ envvars ensure

# create an env file
$ envvars envfile
$ cat .env
# ECHO

# create a declaration file with example value
# envvars:
#   - name: ECHO
#     desc: env var ECHO
#     example: Hello World
$ printf "envvars:\n  - name: ECHO\n    desc: env var ECHO\n    example: Hello World\n" > envvars.yml

# create an env file with the example value
$ envvars envfile -overwrite -example
$ cat .env
# ECHO=Hello World

# explore
$ envvars --help
```

## Principles

Envvars has strict rules which follows some principles.

### Documentation is your best friend

Envvars forces you to have `desc` for your environment variables and tags. This helps anyone new to the project, or juggling with many projects at once, to understand every environment variable, and tag, as long as its `desc` is meaningful.

### You ain't gonna need it

Envvars will complain if a tag is declared but not being used by an environment variable. It will also throw an error if an environment variable uses a tag that is not declared. Lastly, it will not like it if a tag passed as parameter to a CLI command does not exist in the declaration file. All of this helps to prevent issues down the track.

## Feedback

Feedback is greatly appreciated. Do you have workflows that the tool supports well, or doesn't support at all? Do any of the commands have surprising effects, output, or results? Let us know by filing an issue, describing what you did or wanted to do, what you expected to happen, and what actually happened.

## Contributing

Contributions are greatly appreciated. See [CONTRIBUTING.md](https://github.com/flemay/envvars/blob/master/CONTRIBUTING.md) for more details.