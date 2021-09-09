VERSION ?= 0.0.7
GIT_TAG = v$(VERSION)
DOCKER_TAG = $(VERSION)
IMAGE_NAME = flemay/envvars:$(VERSION)
COMPOSE_RUN_GOLANG = docker-compose run --rm golang
COMPOSE_RUN_SHELLCHECK = docker-compose run --rm shellcheck
COMPOSE_RUN_GOLANGCILINT = docker-compose run --rm golangcilint
ENVFILE ?= env.template
TARGET_RUN_ARGS ?= --help

all:
	ENVFILE=env.example $(MAKE) envfile deps test build run buildDockerImage clean

ciTest: envfile deps test build run buildDockerImage clean

_ciRelease:
	TAG=$(GIT_TAG) ./scripts/github_release.sh

envfile:
	cp -f $(ENVFILE) .env

deps:
	docker-compose pull
	$(COMPOSE_RUN_GOLANG) make _deps
_deps:
	go mod download
	go mod vendor

updateDeps:
	$(COMPOSE_RUN_GOLANG) make _updateDeps
_updateDeps:
	go get -d -u ./...
	go mod vendor
	go mod tidy

test:
	$(COMPOSE_RUN_GOLANG) make _test
	$(COMPOSE_RUN_SHELLCHECK) scripts/*.sh
	$(warning [WARNING] golangci-lint - excluding test files should only be temporary until the latest docker image uses go 1.17+)
	$(COMPOSE_RUN_GOLANGCILINT) golangci-lint run --tests=false pkg/...
_test:
	go test -coverprofile=profile.out ./...

build:
	$(COMPOSE_RUN_GOLANG) bash -c 'VERSION=$(VERSION) ./scripts/build.sh'

run:
	$(COMPOSE_RUN_GOLANG) make _run TARGET_RUN_ARGS="$(TARGET_RUN_ARGS)"
_run:
	./bin/envvars $(TARGET_RUN_ARGS)

buildDockerImage:
	docker build --no-cache -t $(IMAGE_NAME) .
	docker run --rm $(IMAGE_NAME) --help
	docker run --rm $(IMAGE_NAME) version

# pushDockerImage pushes the locally built image to docker hub.
# It requires 2 environment variables upfront because it does not use
# a container with access to the Docker docker (which .env could be passed).
pushDockerImage: env-DOCKER_USERNAME env-DOCKER_ACCESS_TOKEN
	@echo "$(DOCKER_ACCESS_TOKEN)" | docker login --username "$(DOCKER_USERNAME)" --password-stdin docker.io
	docker push $(IMAGE_NAME)
	docker logout

removeDockerImage:
	docker rmi -f $(IMAGE_NAME)

tag:
	git tag $(GIT_TAG)
	git push origin $(GIT_TAG)

# this is to be used with caution
overwriteTag:
	-git tag -d $(GIT_TAG)
	-git push origin :refs/tags/$(GIT_TAG)
	git tag $(GIT_TAG)
	git push origin $(GIT_TAG)

clean:
	$(COMPOSE_RUN_GOLANG) make _clean
	docker-compose down --remove-orphans
	-$(MAKE) removeDockerImage
	rm -f .env
_clean:
	rm -fr bin vendor profile.out

shell:
	$(COMPOSE_RUN_GOLANG) bash

env-%:
	$(info Check if $* is not empty)
	@docker run --rm -e ENV_VAR=$($*) node:alpine sh -c '[ -z "$$ENV_VAR" ] && echo "Error: $* is empty" && exit 1 || exit 0'
