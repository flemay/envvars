VERSION = 0.0.1
IMAGE_NAME ?= flemay/envvars:$(VERSION)

deps:
	dep ensure
.PHONY: deps

test:
	go test -cover ./...
.PHONY: test

build:
	go build -o bin/envvars
.PHONY: build

install:
	go install
.PHONY: install

buildForScratch:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/envvars
.PHONY: buildForScratch

dockerBuild:
	docker build --no-cache -t $(IMAGE_NAME) .
.PHONY: dockerBuild

dockerRun:
	docker run --rm $(IMAGE_NAME)
.PHONY: dockerRun

tag:
	-git tag -d $(VERSION)
	-git push origin :refs/tags/$(VERSION)
	git tag $(VERSION)
	git push origin $(VERSION)
.PHONY: tag

clean:
	rm -f bin vendor
.PHONY: clean