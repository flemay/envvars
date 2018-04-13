# Docker Release Process

## Docker Hub Setup

This section is a step-by-step guide on how I configured Docker Hub for building `flemay/envvars` every time a new tag is pushed.

1. Go to [https://hub.docker.com](https://hub.docker.com) and sign in
1. Go to `Create` and  `Create Automated Build`
1. Select Github
1. Select User `flemay` and then the repository `envvars`
1. Fill out the form (Namespace, Name, Visibility, and Short Description)
1. Since the builds will be based on tags only, click on `Click here to customize`
1. There is 1 Autobuild Tag
    1. Push Type: `Tag`
    1. Name: `/.*/`
    1. Dockerfile Location: `/`
    1. Docker Tag: _empty_
1. Click on button `Create`

## Release

> This assumes master contains the latest changes to be released.

1. Update version in Makefile
1. Update version in README.md
1. Build the image locally `$ make dockerBuild`
1. Test the image locally `$ make dockerTest`
1. Commit the changes and push
1. Run `$ make tag`
1. Go to [flemay/envvars](https://hub.docker.com/r/flemay/envvars) on Docker Hub
1. In `Build Details` tab, you should now see the build kicking off
