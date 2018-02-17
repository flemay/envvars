VERSION = 0.0.1
IMAGE_NAME ?= flemay/envvars:$(VERSION)
GOLANG_DEPS_DIR = vendor

deps:
	docker-compose run --rm golang make _deps
.PHONY: deps

test: $(GOLANG_DEPS_DIR)
	docker-compose run --rm golang make _test
.PHONY: test

build: $(GOLANG_DEPS_DIR)
	docker-compose run --rm golang make _build
.PHONY: build

dockerBuild:
	docker build --no-cache -t $(IMAGE_NAME) .
.PHONY: dockerBuild

dockerRun:
	docker run --rm $(IMAGE_NAME)
.PHONY: dockerRun

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

_buildForScratch:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/envvars
.PHONY: _buildForScratch

_build:
	go build -o bin/envvars
.PHONY: _build

_install:
	go install
.PHONY: _install