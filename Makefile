VERSION = 0.0.3
IMAGE_NAME ?= flemay/envvars:$(VERSION)
GOLANG_DEPS_DIR = vendor
EXECUTABLE = bin/envvars
PROFILE_NAME ?= profile.out
COMPOSE_RUN_GOLANG = docker-compose run --rm golang
ENVFILE = .env
DOCKER_RUN_ENVVARS = docker run --rm -v $(PWD):/opt/app -w /opt/app flemay/envvars:$(VERSION)
COMPOSE_RUN_ENVVARS = docker-compose run --rm envvars
COMPOSE_RUN_MUSKETEERS = docker-compose run --rm musketeers

all: envfileExample deps test build run dockerBuild dockerTest clean
.PHONY: all

travis: deps test build run dockerBuild dockerTest triggerDockerHubBuilds clean
.PHONY: travis

.env:
	$(DOCKER_RUN_ENVVARS) envfile

envfileExample:
	$(DOCKER_RUN_ENVVARS) envfile --example --overwrite
.PHONY: envfileExample

deps: $(ENVFILE)
	$(COMPOSE_RUN_GOLANG) make _deps
.PHONY: deps

test: $(ENVFILE) $(GOLANG_DEPS_DIR)
	$(COMPOSE_RUN_GOLANG) make _test
.PHONY: test

build: $(ENVFILE) $(GOLANG_DEPS_DIR)
	$(COMPOSE_RUN_GOLANG) make _build
.PHONY: build

run: $(ENVFILE) $(EXECUTABLE)
	$(COMPOSE_RUN_GOLANG) make _run
.PHONY: run

shell: $(ENVFILE)
	$(COMPOSE_RUN_GOLANG) bash
.PHONY: shell

dockerBuild:
	docker build --no-cache -t $(IMAGE_NAME) .
.PHONY: dockerBuild

dockerTest:
	docker run --rm $(IMAGE_NAME)
	docker run --rm $(IMAGE_NAME) version
.PHONY: dockerTest

dockerRemove:
	docker rmi -f $(IMAGE_NAME)
.PHONY: dockerRemove

tag:
	-git tag -d $(VERSION)
	-git push origin :refs/tags/$(VERSION)
	git tag $(VERSION)
	git push origin $(VERSION)
.PHONY: tag

clean: $(ENVFILE)
	$(COMPOSE_RUN_GOLANG) make _clean
	docker-compose down --remove-orphans
	-$(MAKE) dockerRemove
.PHONY: clean

mock: $(ENVFILE) $(GOLANG_DEPS_DIR)
	$(COMPOSE_RUN_GOLANG) make _mock
.PHONY: mock

triggerDockerHubBuilds: $(ENVFILE)
	$(COMPOSE_RUN_ENVVARS) ensure
	$(COMPOSE_RUN_MUSKETEERS) make _triggerDockerHubBuildOnBranchMasterUpdate
	$(COMPOSE_RUN_MUSKETEERS) make _triggerDockerHubBuildOnGitTagUpdate
.PHONY: triggerDockerHubBuilds

_deps:
	dep ensure
.PHONY: _deps

_test:
	go test -coverprofile=$(PROFILE_NAME) ./...
.PHONY: _test

_build:
	VERSION=$(VERSION) ./scripts/build.sh
.PHONY: _build

_run:
	./$(EXECUTABLE)
.PHONY: _run

_install:
	go install
.PHONY: _install

_mock:
	go get -u github.com/vektra/mockery/.../
	mockery -dir=pkg -all -case=underscore -output=pkg/mocks
.PHONY: _mock

_htmlCover:
	go tool cover -html=$(PROFILE_NAME)
.PHONY: _htmlCover

_clean:
	rm -fr bin vendor
.PHONY: _clean

_triggerDockerHubBuildOnBranchMasterUpdate:
	[ "$(TRAVIS_BRANCH)" = "master" ] \
	&& curl -H "Content-Type: application/json" --data '{"docker_tag": "latest"}' -X POST $(DOCKERHUB_TRIGGER_URL) \
	&& echo "TRIGGERED Docker build for branch master" \
	|| echo "SKIPPED Docker build for branch master"
.PHONY: _triggerDockerHubBuildOnBranchMasterUpdate

_triggerDockerHubBuildOnGitTagUpdate:
	[ "$(TRAVIS_BRANCH)" != "master" ] \
	&& [ -n "$(TRAVIS_TAG)" ] \
	&& [ "$(TRAVIS_PULL_REQUEST)" = "false" ] \
	&& curl --data '{"source_type": "Tag", "source_name": "$(TRAVIS_TAG)"}' -X POST $(DOCKERHUB_TRIGGER_URL) \
	&& echo "TRIGGERED Docker build for tag $(TRAVIS_TAG)" \
	|| echo "SKIPPED Docker builds for tag"
.PHONY: _triggerDockerHubBuildOnGitTagUpdate
