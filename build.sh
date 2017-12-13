#!/usr/bin/env bash
export GOROOT=/home/go
export GOBIN="${GOROOT}/bin/go"
export GOPATH=`pwd`
cd src/main
$GOBIN build
cd ../..
cp src/main/main bin/main
echo "build over"

