#!/bin/bash

set -e

pushd system_tests
	bundle
	bundle exec rspec
popd

exit 0
