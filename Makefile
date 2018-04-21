VERSION = 0.0.2
IMAGE_NAME ?= flemay/envvars:$(VERSION)
GOLANG_DEPS_DIR = vendor
EXECUTABLE = bin/envvars

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
	rm -fr bin vendor
.PHONY: clean

_deps:
	dep ensure
.PHONY: _deps

_test:
	go test -cover ./...
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