VERSION ?= 0.0.5
GIT_TAG = v$(VERSION)
DOCKER_TAG = $(VERSION)
IMAGE_NAME = flemay/envvars:$(VERSION)
GOLANG_DEPS_DIR = vendor
EXECUTABLE = bin/envvars
PROFILE_NAME = profile.out
COMPOSE_RUN_GOLANG = docker-compose run --rm golang
ENVFILE ?= env.template

all:
	ENVFILE=env.example $(MAKE) envfile deps test build run buildDockerImage testDockerImage clean

travis:
	GIT_TAG=$(GIT_TAG) ./scripts/travis.sh

onPullRequest: all

onMasterChange: envfile deps test sendCoverage build run buildDockerImage testDockerImage clean

onGitTag: envfile deps test build run buildDockerImage testDockerImage pushDockerImage clean

envfile:
	cp -f $(ENVFILE) .env

deps:
	$(COMPOSE_RUN_GOLANG) make _deps

_deps:
	go mod download
	go mod vendor

test: $(GOLANG_DEPS_DIR)
	$(COMPOSE_RUN_GOLANG) make _test

_test:
	go test -coverprofile=$(PROFILE_NAME) ./...

build: $(GOLANG_DEPS_DIR)
	$(COMPOSE_RUN_GOLANG) make _build

_build:
	VERSION=$(VERSION) ./scripts/build.sh

run: $(EXECUTABLE)
	$(COMPOSE_RUN_GOLANG) make _run

_run:
	./$(EXECUTABLE)

shell:
	$(COMPOSE_RUN_GOLANG) bash

sendCoverage: $(PROFILE_NAME)
	$(COMPOSE_RUN_GOLANG) bash -c 'bash <(curl -s https://codecov.io/bash) -f $(PROFILE_NAME)'

clean:
	$(COMPOSE_RUN_GOLANG) make _clean
	docker-compose down --remove-orphans
	-$(MAKE) removeDockerImage

_clean:
	rm -fr bin vendor

_tag:
	-git tag -d $(GIT_TAG)
	-git push origin :refs/tags/$(GIT_TAG)
	git tag $(GIT_TAG)
	git push origin $(GIT_TAG)

################
# DOCKER IMAGE #
################

buildDockerImage:
	docker build --no-cache -t $(IMAGE_NAME) .

testDockerImage:
	docker run --rm $(IMAGE_NAME)
	docker run --rm $(IMAGE_NAME) version

pushDockerImage:
	./scripts/push.sh

removeDockerImage:
	docker rmi -f $(IMAGE_NAME)

###########
# MOCKERY #
###########

mock: $(GOLANG_DEPS_DIR)
	$(COMPOSE_RUN_GOLANG) make _mock

_mock:
	go get -u github.com/vektra/mockery/.../
	mockery -dir=pkg -all -case=underscore -output=pkg/mocks