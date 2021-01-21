#!/bin/bash

if [ "$CC" = "clang-11" ]; then
   wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key|sudo apt-key add -
   sudo apt-add-repository 'deb http://apt.llvm.org/focal/ llvm-toolchain-focal-11 main'
   sudo apt-get install -y llvm-11 clang-11
fi

eval "$(gimme 1.15.4)"
source ~/.gimme/envs/go1.15.4.env

export GOPATH=~/go:$(pwd)/build/generated/golang
export PATH=$PATH:~/go/bin

go get -u golang.org/x/lint/golint
go get github.com/btcsuite/btcutil/base58

if [ "$RUN_TYPE" = "coverage" ]; then
   sudo apt-get install -y lcov ruby
   sudo gem install coveralls-lcov
   go get -u github.com/jandelgado/gcov2lcov
fi

pip3 install --user dataclasses-json Jinja2 importlib_resources pluginbase gitpython
