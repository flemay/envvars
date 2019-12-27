#!/usr/bin/env bash
set -e
set -u

go get -u github.com/vektra/mockery/.../
mockery -dir=pkg -all -case=underscore -output=pkg/mocks
