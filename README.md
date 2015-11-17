# GoPretend

A Golang Service Virtualization Framework

## To download

```
go get github.com/dhrapson/gopretend
```
This will download and install the go executable, then put it into your $GOPATH/bin directory.
It would be a good idea to add $GOPATH/bin to your PATH

## To run

```
gopretend [path/to/file.yml]
```
The gopretend executable will use a the config file you provide it.
If no file is provided, it will look for a gopretend.yml file in the current directory.
If no local gopretend.yml file is found, the gopretend service will start and await configuration via API.

## To execute system tests

```
cd system_tests
bundle
bundle exec rspec
```
