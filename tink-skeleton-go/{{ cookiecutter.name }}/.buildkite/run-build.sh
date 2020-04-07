#!/bin/bash
set -euo pipefail

export CGO_ENABLED=0
go generate ./...
go build ./cmd/{{ cookiecutter.name }}
mkdir -p bin
mv ./{{ cookiecutter.name }} bin/
