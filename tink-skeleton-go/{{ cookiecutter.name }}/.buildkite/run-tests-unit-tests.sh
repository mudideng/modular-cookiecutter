#!/bin/bash
set -eo pipefail

# Running one test file ata a time to ignore vendor directory and collect
# coverage data. See [1].
#
# [1] https://github.com/codecov/example-go#caveat-multiple-files
echo "" > coverage.txt

for d in $(go list ./... | grep -v vendor); do
    go test -v -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done

if [[ -z "$CODECOV_TOKEN" ]];then
    echo Not uploading coverage data because '$CODECOV_TOKEN' was not set.
else
    echo Uploading coverage data...
    bash <(curl -s https://codecov.io/bash) -F unittests
fi
