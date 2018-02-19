# Contributing to Envvars

Envvars is an open source project and contributions are greatly appreciated. Before starting any work, please either comment on an existing issue, or file a new one.

## Contributing code

The project follows the typical GitHub pull request model.

### 1. Fork

Create a fork of Envvars on GitHub so that you can commit your changes and create a Pull Request.

### 2. Development

There are 2 approaches to make your changes:

- [3 Musketeers](https://github.com/flemay/3musketeers) which uses Make, Docker, and Compose
- a Go environment with dep

#### 3 Musketeers

##### Prerequisites

- Docker
- Compose
- Make

##### Steps

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

#### Go environment

##### Prerequisites

- Go 1.10
- [dep](https://github.com/golang/dep)
- Make

##### Steps

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

### 3. Create a Pull Request