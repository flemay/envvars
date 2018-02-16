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
