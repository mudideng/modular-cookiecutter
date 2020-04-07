{{ cookiecutter.name }}
===========
Go service created from github.com/tink-ab/tink-skeleton-go

Package structure adheres to [Ben Johnson's Standard Package
Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1).


What to do after forking this repository?
-----------------------------------------

1. Create a pipeline in Buildkite.
1. Adjust the README to your project.
1. Write a new microservice!


Developing
----------
Install Go development environment. Clone this repository:

```bash
$ git clone git@github.com:tink-ab/{{ cookiecutter.name }}.git $GOPATH/src/github.com/tink-ab/{{ cookiecutter.name }}
```

You should be able to develop using basic Golang tooling.

### Dependencies and building

```bash
$ go build ./cmd/{{ cookiecutter.name }}/{{ cookiecutter.name }}.go
$ # Test running this application (but don't forget to set the correct command line arguments):
$ ./{{ cookiecutter.name }}
```

### Running the application locally

And then you can start the service with

```bash
go run cmd/{{ cookiecutter.name }}/{{ cookiecutter.name }}.go
```
. You can list the available subcommands using `go run cmd/{{ cookiecutter.name }}/{{ cookiecutter.name }}.go --help`.

### Running them tests

#### Unit tests

You run the tests by executing

```bash
$ go test -v ./...
```

#### Goconvey

For running the tests continuously while development, have a look at
[Goconvey](http://goconvey.co):

```bash
$ goconvey -timeout 20s -excludedDirs vendor
```
or
```bash
$ CASSANDRA_HOST=172.16.42.2 MYSQL_HOST=172.16.42.2 goconvey -timeout 20s -excludedDirs vendor
```
to also execute integration and (possibly) system tests.

### Reading documentation

It can be beneficial to look at the documentation to get the know the code.
You can browse the documentation by running `godoc -http=:6060` and opening
up [http://localhost:6060](http://localhost:6060).
