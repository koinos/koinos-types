#!/bin/bash

set -e
set -x

if [ "$RUN_TYPE" = "test" ]; then
   # C++ tests
   cd $(dirname "$0")/../build/tests
   exec ctest -j3 --output-on-failure

   # Golang tests
   cd $(dirname "$0")/..
   GOPATH=$GOPATH:$(pwd)/build/generated/golang
   go get -d ./tests/golang
   go test ./tests/golang
fi

