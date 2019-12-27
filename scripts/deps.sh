#!/usr/bin/env bash
set -e
set -u

go mod download
go mod vendor