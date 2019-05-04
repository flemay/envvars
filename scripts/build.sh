#!/usr/bin/env bash
set -e
set -u

if [ -z ${VERSION+x} ]; then echo "VERSION is not set"; exit 1; fi

COMMIT_HASH=$(git rev-parse --short HEAD 2>/dev/null)
CURRENT_DATE=$(date "+%Y-%m-%d")

GO_BUILD_LDFLAGS="-s -w -X main.commitHash=$COMMIT_HASH -X main.buildDate=$CURRENT_DATE -X main.version=$VERSION"

if [ "$IS_SCRATCH_IMAGE" = true ]; then
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "$GO_BUILD_LDFLAGS" -o bin/envvars
else
  go build -ldflags "$GO_BUILD_LDFLAGS" -o bin/envvars
fi