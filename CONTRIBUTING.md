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

# checkout a new branch if you do not want to work on master(optional)
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
$ git push fork master
# or your branch
$ git push fork meaningful_branch_name

# create pull request https://help.github.com/articles/creating-a-pull-request/
```

> Steps from [Francesc's tweet](https://mobile.twitter.com/francesc/status/1009487969198075905)

### Mocks

[Mockery](https://github.com/vektra/mockery) is being used to generate mocks in the folder `pkg/mocks`.

- Mockery must be installed locally
- To generate all of them `$ make _mock`
- Once generated, one would need to fix all the mocks imports