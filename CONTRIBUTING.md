# Contributing to Envvars

Envvars is an open source project and contributions are greatly appreciated.

There are different ways to contribute to this project.

## Star it

If you like the project, go on [GitHub](https://github.com/flemay/envvars) and ⭐️ it!

## Share it

If you think the 3 Musketeers would be valuable to your friends, teammates, and company, share it!

## Feedback

Feedback is greatly appreciated. Do you have workflows that the tool supports well, or doesn't support at all? Do any of the commands have surprising effects, output, or results? Let us know by filing an issue, describing what you did or wanted to do, what you expected to happen, and what actually happened.

## Contributing code

The project follows the typical GitHub pull request model. Before starting any work, please either comment on an existing issue, or file a new one.

envvars development uses the [3 Musketeers pattern](https://github.com/flemay/3musketeers) which requires Docker, Compose, and Make. Any change must be releasable.

```bash
# fork https://github.com/flemay/envvars -> github.com/my-id/project
# get the project
$ go get github.com/flemay/envvars
# go inside the repo
$ cd $GOPATH/src/github.com/flemay/envvars && git remote add fork git@github.com:my-id/project.git

# checkout a new branch if you do not want to work on main (optional)
$ git checkout -b meaningful_branch_name

# download all the dependencies
$ make deps
# make your changes
# ...
# test your changes
$ make test
# build envvars
$ make build
# run envvars
$ make run
# build a local docker image
$ make buildDockerImage

# push your changes
$ git push fork main
# or your branch
$ git push fork meaningful_branch_name

# create pull request https://help.github.com/articles/creating-a-pull-request/
```

> Steps from [Francesc's tweet](https://mobile.twitter.com/francesc/status/1009487969198075905)

### cmd/envvars/version.json

This file contains information that are usually provided during the build process. This file is special as it requires it when developing but further modifications of the file (like during the build process)  are not commited. Git `--skip-worktree` is used. Credits to https://compiledsuccessfully.dev/git-skip-worktree/.

Steps for updating the file

1. `$ git update-index --no-skip-worktree cmd/envvars/version.json`
1. Commit the changes and push
1. `$ git update-index --skip-worktree cmd/envvars/version.json`
1. List files that are skipped: `$ git ls-files -v | grep '^S'`
    1. `version.json` should be included

Steps for fixing an error like `error: Your local changes to the following files would be overwritten by checkout: path/to/file` when checking out a different branch

1. `$ git update-index --no-skip-worktree cmd/envvars/version.json`.
1. `$ git stash` (or maybe `$ git checkout cmd/envvars/version.json`)
1. `$ git checkout <your-branch>`
1. `$ git update-index --skip-worktree cmd/envvars/version.json`
1. Can now clean the stash

### Go modules

To update modules, run the command `$ make updateDeps`.
