# Resemble

A declaritive Service Virtualization and Stubbing utility written in Golang

This software is currently under contruction, prior to initial release, using the Product Model style of management.
The publicly available [product backlog](https://www.pivotaltracker.com/n/projects/1485132) can give you some idea of the intended direction and velocity.

## Intended users
The intention is to provide an easy way to configure a lightweight HTTP service that imitates a real-world HTTP service of some kind. Using resemble would mean that a full copy of the HTTP service being consumed might not be required until staging / production. Typical usage patterns will include:

* a developer running resemble on his/her laptop to create an imitation web service locally, in support of UI testing
* a developer running resemble to declare a stub in support of integration testing
* a CI server starting resemble to run some immitation services in support of integration / system tests
* a CI server or perf tester running some immitation services that are configured to run a certain performance thresheold (e.g. no more than 1 TPS)
* a CI server or functional tester running some immitation services that are configured to fail in a realistic way in order to test failure cases in a consumer app
* a developer wanting to run a simple HTTP static file server locally without the need for a heavy installation

## To download

```
go get github.com/dhrapson/resemble/resemble
```
This will download and install the go executable, then put it into your $GOPATH/bin directory.
It would be a good idea to add $GOPATH/bin to your PATH

## To run

```
resemble [path/to/file.yml]
```
The resemble executable will use a the config file you provide it.
If no file is provided, it will look for a resemble.yml file in the current directory.
If no local resemble.yml file is found, the resemble service will start and await configuration via API.

## To execute system tests

```
cd system_tests
bundle
bundle exec rspec
```
