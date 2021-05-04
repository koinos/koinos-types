#!/bin/bash

sudo apt-get install -yq --allow-downgrades libc6=2.31-0ubuntu9.2 libc6-dev=2.31-0ubuntu9.2
sudo -E apt-get -yq --no-install-suggests --no-install-recommends --allow-downgrades --allow-remove-essential --allow-change-held-packages install clang-11 llvm-11 -o Debug::pkgProblemResolver=yes

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
