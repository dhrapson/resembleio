#!/bin/bash

: "${GOARCH:?Need to set GOARCH non-empty}"
: "${GOOS:?Need to set GOOS non-empty}"

BINDIR="binaries/latest/$GOARCH/$GOOS"
mkdir -p $BINDIR

set -e

pushd resemble
	GOARCH=$GOARCH GOOS=$GOOS go build *.go
popd

BINFILES=`find resemble/ -maxdepth 1 -type f -not -name "*.go"`

echo "BINFILES: $BINFILES"
if [ -z "$BINFILES" ]; then
 	echo "No binary files found to copy for ${GOOS}-${GOARCH}"
	exit 1
fi

for BINFILE in `echo $BINFILES`
do
	echo "Copying $BINFILE to $BINDIR for publishing"
	cp $BINFILE $BINDIR
done
