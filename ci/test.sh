#!/bin/bash

set -e
set -x

if [ "$RUN_TYPE" = "test" ]; then
   # C++ tests
   pushd build/tests/cpp
   ctest -j3 --output-on-failure
   popd

   # Golang tests
   GOPATH=~/go:$(pwd)/build/generated/golang
   go get -d ./tests/golang
   go test ./tests/golang
fi

