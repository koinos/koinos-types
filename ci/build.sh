#!/bin/bash

set -e
set -x

mkdir build
pushd build

if [ "$RUN_TYPE" = "test" ]; then
   if [ "$TRAVIS_OS_NAME" = "osx" ]; then
      cmake -DCMAKE_BUILD_TYPE=Release -GXcode ..
      cmake --build . --config Release --parallel 3 -- -quiet
   else
      cmake -DCMAKE_BUILD_TYPE=Release ..
      cmake --build . --config Release --parallel 3
   fi
elif [ "$RUN_TYPE" = "coverage" ]; then
   cmake -DCMAKE_BUILD_TYPE=Debug -DCOVERAGE=ON ..
   cmake --build . --config Debug --parallel 3 --target coverage --target golang

   popd
   go test ./tests/golang -coverprofile=./build/go-coverage.out -coverpkg=./build/generated/golang/src/koinos
   gcov2lcov -infile=./build/go-coverage.out -outfile=./build/go-coverage.info

   lcov -a ./build/coverage.info -a ./build/go-coverage.info -o ./build/merged.info
fi

