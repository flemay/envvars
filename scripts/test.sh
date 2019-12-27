#!/usr/bin/env bash
set -e
set -u

go test -coverprofile=profile.out ./...