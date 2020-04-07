FROM gcr.io/tink-containers/alpine:3.8

RUN apk add --no-cache ca-certificates

COPY {{ cookiecutter.name }} /

ENTRYPOINT ["/{{ cookiecutter.name }}", "serve"]
