deps:
	dep ensure
.PHONY: deps

test:
	go test -cover ./...
.PHONY: test

install:
	go install
.PHONY: install