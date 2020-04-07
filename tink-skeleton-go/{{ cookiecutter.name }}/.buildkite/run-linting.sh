#!/bin/bash
set -euo pipefail

# golangci-lint doesn't seem to support Go modules so we create a vendor
# directory and use that.
go mod vendor
export GO111MODULE=off
golangci-lint run
