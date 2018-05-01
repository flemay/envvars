VERSION = 0.0.3
IMAGE_NAME ?= flemay/envvars:$(VERSION)
GOLANG_DEPS_DIR = vendor
EXECUTABLE = bin/envvars
PROFILE_NAME ?= profile.out

deps:
	docker-compose run --rm golang make _deps
.PHONY: deps

test: $(GOLANG_DEPS_DIR)
	docker-compose run --rm golang make _test
.PHONY: test

build: $(GOLANG_DEPS_DIR)
	docker-compose run --rm golang make _build
.PHONY: build

run: $(EXECUTABLE)
	docker-compose run --rm golang make _run
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
	docker-compose down --remove-orphans
	rm -fr bin vendor
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