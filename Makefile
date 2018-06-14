VERSION = 0.0.4
GIT_TAG = v$(VERSION)
DOCKER_TAG = $(VERSION)
IMAGE_NAME = flemay/envvars:$(VERSION)
GOLANG_DEPS_DIR = vendor
EXECUTABLE = bin/envvars
PROFILE_NAME = profile.out
COMPOSE_RUN_GOLANG = docker-compose run --rm golang
DOCKER_RUN_ENVVARS = docker run --rm -v $(PWD):/opt/app -w /opt/app flemay/envvars:$(VERSION)
COMPOSE_RUN_ENVVARS = docker-compose run --rm envvars

# all shows how the pipeline looks like.
# dockerBuild dockerTest are first because other tasks will need a .env and
# this is required when testing a new version which does not exist yet in
# Docker Hub.
all: envfileExample dockerBuild dockerTest deps test build run clean
.PHONY: all

# travis is used by Travis CI for its build.
travis: envfile dockerBuild dockerTest triggerDockerHubBuilds deps test sendCoverage clean
.PHONY: travis

# envfile creates a .env with envvars unless ENVFILE is defined, in which case
# copies the file to .env
envfile:
	[ -n "$(ENVFILE)" ] && cp -f $(ENVFILE) .env || $(DOCKER_RUN_ENVVARS) envfile --overwrite
.PHONY: envfile

envfileExample:
	$(DOCKER_RUN_ENVVARS) envfile --example --overwrite
.PHONY: envfileExample

deps:
	$(COMPOSE_RUN_GOLANG) make _deps
.PHONY: deps

test: $(GOLANG_DEPS_DIR)
	$(COMPOSE_RUN_GOLANG) make _test
.PHONY: test

build: $(GOLANG_DEPS_DIR)
	$(COMPOSE_RUN_GOLANG) make _build
.PHONY: build

run: $(EXECUTABLE)
	$(COMPOSE_RUN_GOLANG) make _run
.PHONY: run

shell:
	$(COMPOSE_RUN_GOLANG) bash
.PHONY: shell

sendCoverage: $(PROFILE_NAME)
	$(COMPOSE_RUN_GOLANG) bash -c 'bash <(curl -s https://codecov.io/bash) -f $(PROFILE_NAME)'
.PHONY: sendCoverage

clean:
	$(COMPOSE_RUN_GOLANG) make _clean
	docker-compose down --remove-orphans
	-$(MAKE) dockerRemove
.PHONY: clean

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

_htmlCover:
	go tool cover -html=$(PROFILE_NAME)
.PHONY: _htmlCover

_clean:
	rm -fr bin vendor
.PHONY: _clean

_tag:
	-git tag -d $(GIT_TAG)
	-git push origin :refs/tags/$(GIT_TAG)
	git tag $(GIT_TAG)
	git push origin $(GIT_TAG)
.PHONY: _tag

################
# DOCKER IMAGE #
################

# dockerBuild always builds the Docker image with no cache.
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

###########
# MOCKERY #
###########

mock: $(GOLANG_DEPS_DIR)
	$(COMPOSE_RUN_GOLANG) make _mock
.PHONY: mock

_mock:
	go get -u github.com/vektra/mockery/.../
	mockery -dir=pkg -all -case=underscore -output=pkg/mocks
.PHONY: _mock

#######################
# DOCKER HUB TRIGGERS #
#######################

triggerDockerHubBuilds:
	$(COMPOSE_RUN_ENVVARS) ensure
	$(COMPOSE_RUN_GOLANG) make _triggerDockerHubLatestBuildOnBranchMasterUpdate \
		_triggerDockerHubTagBuildOnGitTagUpdate \
		_triggerDockerHubAllBuildsIfCronJob
.PHONY: triggerDockerHubBuilds

_triggerDockerHubLatestBuildOnBranchMasterUpdate:
	@if [ "$(TRAVIS_BRANCH)" = "master" ]; then \
		curl -H "Content-Type: application/json" --data '{"docker_tag": "latest"}' -X POST $(DOCKERHUB_TRIGGER_URL); \
		echo " TRIGGERED Docker build for branch master"; \
	else
		echo " SKIPPED Docker build for branch master"; \
	fi;
.PHONY: _triggerDockerHubLatestBuildOnBranchMasterUpdate

_triggerDockerHubTagBuildOnGitTagUpdate:
	@if [ "$(TRAVIS_BRANCH)" != "master" ] && [ -n "$(TRAVIS_TAG)" ]; then \
		if [ "$(TRAVIS_TAG)" != "$(GIT_TAG)" ]; then \
			echo "TRAVIS_TAG ($(TRAVIS_TAG)) cannot be different than GIT_TAG ($(GIT_TAG))"; \
			exit 1; \
		fi; \
		curl -H "Content-Type: application/json" --data '{"source_type": "Tag", "source_name": "$(DOCKER_TAG)"}' -X POST $(DOCKERHUB_TRIGGER_URL); \
		echo " TRIGGERED Docker build for tag $(DOCKER_TAG)"; \
	else \
		echo " SKIPPED Docker builds for tag"; \
	fi;
.PHONY: _triggerDockerHubTagBuildOnGitTagUpdate

_triggerDockerHubAllBuildsIfCronJob:
	@if [ "$(TRAVIS_EVENT_TYPE)" = "cron" ] && [ "$(TRAVIS_PULL_REQUEST)" = "false" ] && [ "$(TRAVIS_BRANCH)" = "master" ]; then \
		curl -H "Content-Type: application/json" --data '{"build": true}' -X POST $(DOCKERHUB_TRIGGER_URL); \
		echo " TRIGGERED Docker builds for all tags"; \
	else \
		echo " SKIPPED Docker builds for all tags"; \
	fi;
.PHONY: _triggerDockerHubAllBuildsIfCronJob
