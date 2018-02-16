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

dockerImage:
	docker build --no-cache -t flemay/envvars .
.PHONY: dockerImage