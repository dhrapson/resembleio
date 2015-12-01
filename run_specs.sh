#!/bin/bash

PATH=$PATH:$GOPATH/bin

echo $PWD
pushd `dirname $0`
	set -e

	pushd resemble
		go install
	popd

	pushd system_tests
		bundle
		bundle exec rspec
	popd
popd

exit 0
