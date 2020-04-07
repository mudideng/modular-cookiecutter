FROM gcr.io/tink-containers/golang:1.11

# This will install the golangci-lint binary in ./bin/
RUN cd $GOPATH && curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s v1.11.2

ENV GO111MODULE=on

RUN mkdir /root/.ssh/ && ssh-keyscan -t rsa github.com >> /root/.ssh/known_hosts

COPY gitconfig /root/.gitconfig
