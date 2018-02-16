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
	docker build --no-cache -t flemay/envvars .
.PHONY: dockerBuild

dockerRun:
	docker run --rm flemay/envvars
.PHONY: dockerRun