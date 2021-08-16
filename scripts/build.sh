#!/usr/bin/env bash
set -e
set -u

if [ -z ${VERSION+x} ]; then echo "VERSION is not set"; exit 1; fi

COMMIT_HASH=$(git rev-parse --short HEAD 2>/dev/null)
CURRENT_DATE=$(date "+%Y-%m-%d")
VERSION_JSON_FILE="
{
  \"Version\": \"$VERSION\",
  \"BuildDate\": \"$CURRENT_DATE\",
  \"GitCommit\": \"$COMMIT_HASH\"
}
"
echo $VERSION_JSON_FILE > cmd/envvars/version.json

# -s -w: omit symbol table, debug information, and DWARF table
GO_BUILD_LDFLAGS="-s -w -extldflags \"-static\""
if [ "$BUILD_FOR_SCRATCH_IMAGE" = true ]; then
  GOOS=linux GOARCH=amd64 go build -ldflags "$GO_BUILD_LDFLAGS" -o bin/envvars cmd/envvars/main.go
else
  go build -ldflags "$GO_BUILD_LDFLAGS" -o bin/envvars
fi
