#!/bin/bash

set -e
set -x

mkdir build
pushd build

if [ "$RUN_TYPE" = "test" ]; then
   cmake -DCMAKE_BUILD_TYPE=Release ..
   cmake --build . --config Release --parallel 3
elif [ "$RUN_TYPE" = "coverage" ]; then
   cmake -DCMAKE_BUILD_TYPE=Debug -DCOVERAGE=ON ..
   cmake --build . --config Debug --parallel 3 --target coverage --target golang

   popd
   go test ./tests/golang -coverprofile=./build/go-coverage.out -coverpkg=./build/generated/golang/src/github.com/koinos/koinos-types-golang
   go test ./build/... -coverprofile=./build/go-generated-coverage.out -coverpkg=./build/generated/golang/src/github.com/koinos/koinos-types-golang
   gcov2lcov -infile=./build/go-coverage.out -outfile=./build/go-coverage.info
   gcov2lcov -infile=./build/go-generated-coverage.out -outfile=./build/go-generated-coverage.info

   lcov -a ./build/coverage.info -a ./build/go-coverage.info -o ./build/merged.info
   lcov -a ./build/merged.info -a ./build/go-generated-coverage.info -o ./build/merged.info
fi
