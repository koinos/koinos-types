#!/bin/bash

set -e
set -x

if [ "$RUN_TYPE" = "test" ]; then
   # C++ tests
   pushd build/tests/cpp
   ctest -j3 --output-on-failure
   popd

   # Golang tests
   go get -d ./build/generated/golang/src/koinos
   go test -v ./tests/golang
fi

