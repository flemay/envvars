# Envvars

Envvars gives the environment variables the love they deserve.

It documents the environment variables of a project to help the team to understand them. It makes sure they are defined before testing, building, deploying, and running. It can also generates an env file to be used by other applications such as Docker and Compose.

## Installation

```bash
# with go get
$ go get -u github.com/flemay/envvars

# or use the tiny docker image (< 5 MB)
$ docker run --rm flemay/envvars:0.0.1
```

## Usage

```bash
# create a declaration file envvars.toml
# [[envvars]]
#   name="ECHO"
#   desc="env var ECHO"
$ printf "[[envvars]]\n  name=\"ECHO\"\n  desc=\"env var ECHO\"" > envvars.toml

# validate the declaration file if it contains errors
$ envvars validate

# ensure the environment variables comply with the declaration file
$ envvars ensure
# Error: environment variable ECHO must be set
# set ECHO
$ export ECHO="helloworld"
$ envvars ensure

# create an env file
$ envvars envfile
$ cat .env
# ECHO

# explore
$ envvars --help
```

## Declaration File

The declaration file (written in [TOML](https://github.com/toml-lang/toml)) is the core of Envvars. It declares all the environment variables used by a project.

Envvars is looking for the declaration file `envvars.toml` by default. A different file can be passed with the flag `-f path/to/declarationfile.toml`.

```toml
# [[tags]] declares a tag.
# They are optional but an error will be thrown if
#  - a [[tags]] is declared but no [[envvars]] uses it
#  - an [[envvars]] uses a tag that is not declared with [[tags]]
[[tags]]
  # name of the tag
  name="deploy"
  # description of the tag
  desc="tag used when deploying"

# [[envvars]] declares an environment variable
[[envvars]]
  # name of the environment variable
  name="ENV"
  # description of the environment variable
  desc="Application's stage (dev, qa, preprod, prod)"
  # tags of the environment variable
  # they are optional but if present, must be declared in [[tags]]
  tags=["deploy"]

[[envvars]]
  name="ENVVAR_1"
  desc="description of ENVVAR_1"
```

## Principles

Envvars has strict rules which follows some principles.

### Documentation is your best friend

Envvars forces you to have `desc` for `[[envvars]]` and `[[tags]]`. This helps anyone new to the project, or juggling with many projects at once, to understand every environment variable, and tag, as long as its `desc` is meaningful.

### You ain't gonna need it

Envvars will complain if `[[tags]]` is declared but not being used by `[[envvars]]`. It will also throw an error if an `[[envvars]]` uses a tag that is not declared. Lastly, it will not like it if a tag passed as parameter to a command does not exist in the declaration file. All of this helps to prevent issues down the track.

## Feedback

Feedback is greatly appreciated. At this stage, the maintainers are most interested in feedback centered on the user experience (UX) of the tool. Do you have workflows that the tool supports well, or doesn't support at all? Do any of the commands have surprising effects, output, or results? Let us know by filing an issue, describing what you did or wanted to do, what you expected to happen, and what actually happened.

## Contributing

Contributions are greatly appreciated. See [CONTRIBUTING.md](http://github.com/flemay/envvars/CONTRIBUTING.md) for more details.