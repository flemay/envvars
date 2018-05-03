VERSION = 0.0.3
IMAGE_NAME ?= flemay/envvars:$(VERSION)
GOLANG_DEPS_DIR = vendor
EXECUTABLE = bin/envvars
PROFILE_NAME ?= profile.out
COMPOSE_RUN_GOLANG = docker-compose run --rm golang

all: clean deps test build run
.PHONY: all

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

clean:
	$(COMPOSE_RUN_GOLANG) make _clean
	docker-compose down --remove-orphans
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

_mock:
	mockery -dir=pkg -all -case=underscore -output=pkg/mocks
.PHONY: _mock

_htmlCover:
	go tool cover -html=$(PROFILE_NAME)
.PHONY: _htmlCover

_clean:
	rm -fr bin vendor
.PHONY: _clean