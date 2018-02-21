# Contributing to Envvars

Envvars is an open source project and contributions are greatly appreciated.

There are few ways to contribute to this project.

## Star it

If you like the project, go on [GitHub](https://github.com/flemay/envvars) and ⭐️ it!

## Share it

If you think Envvars can benefit your friends, teammates, and company, share it!

## Feedback

Feedback is greatly appreciated. Do you have workflows that the tool supports well, or doesn't support at all? Do any of the commands have surprising effects, output, or results? Let us know by filing an issue, describing what you did or wanted to do, what you expected to happen, and what actually happened.

## Contributing code

The project follows the typical GitHub pull request model. Before starting any work, please either comment on an existing issue, or file a new one.

### 1. Fork

[Fork Envvars](https://github.com/flemay/envvars/fork) on GitHub so that you can commit your changes and create a Pull Request.

### 2. Development

There are 2 approaches to make your changes:

- [3 Musketeers](https://github.com/flemay/3musketeers) which uses Make, Docker, and Compose
- a Go environment with dep

#### Development with the 3 Musketeers

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

#### Development with Go environment

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

Happy with all your precious changes? Create a [Pull Request](https://help.github.com/articles/creating-a-pull-request/)!