# envvars

envvars gives the environment variables the love they deserve.

## Installation

```bash
# with go get
$ go get github.com/flemay/envvars

# or use the very tiny (4.51MB) docker image
$ docker run --rm flemay/envvars:0.0.1
```

## Usage

```bash
# create a definition file envvars.toml
# [[envvars]]
#   name="ECHO"
#   desc="env var ECHO"
$ printf "[[envvars]]\n  name=\"ECHO\"\n  desc=\"env var ECHO\"" > envvars.toml

# validate the definition file if it contains errors
$ envvars validate

# ensure the environment variables comply with the definition file
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

## Definition File

The Definition File (written in [TOML](https://github.com/toml-lang/toml)) is the core of envvars. It defines all the environment variables of a given project.

envvars is looking for `envvars.toml` by default but a different file can be passed with the flag `-f path/to/definitionfile.toml`.

```toml
# This is just an example to illustrate what a Definition File looks like

# [[tags]] defines a tag.
# They are optional but an error will be thrown if
#  - a [[tags]] is defined but no [[envvars]] uses it
#  - an [[envvars]] uses a tag that is not defined with [[tags]]
[[tags]]
  # name of the tag
  name="deploy"
  # description of the tag
  desc="tag used when deploying"

# [[envvars]] defines an environment variable
[[envvars]]
  # name of the environment variable
  name="ENV"
  # description of the environment variable
  desc="Application's stage (dev, qa, preprod, prod)"
  # tags of the environment variable
  # they are optional but if present, must be defined in [[tags]]
  tags=["deploy"]

[[envvars]]
  name="ENVVAR_1"
  desc="description of ENVVAR_1"
```