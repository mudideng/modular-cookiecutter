#!/bin/sh

cookiecutter . -f --no-input git_init=0

tar czvf tink-demo-service.tar.gz tink-demo-service
