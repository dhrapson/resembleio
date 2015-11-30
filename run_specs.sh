#!/bin/bash

echo $PWD
pushd `dirname $0`
	set -e

	pushd system_tests
		bundle
		bundle exec rspec
	popd
popd

exit 0
