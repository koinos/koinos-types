#!/bin/bash

set -e
set -x

if [ "$RUN_TYPE" = "test" ]; then
   # C++ tests
   pushd ../build/tests/cpp
   ctest -j3 --output-on-failure
   popd

   # Golang tests
   pushd ..
   GOPATH=$GOPATH:$(pwd)/build/generated/golang
   go get -d ./tests/golang
   go test ./tests/golang
   popd
fi

