#!/bin/bash

if [ "$TRAVIS_OS_NAME" = "linux" ]; then
   sudo apt-get install -y \
      libboost-all-dev \
      python3 \
      python3-pip \
      python3-setuptools \
      valgrind \
      ccache

   wget https://dl.google.com/go/go1.15.3.linux-amd64.tar.gz
   tar -xvf go1.15.3.linux-amd64.tar.gz
   sudo mv go /usr/local
elif [ "$TRAVIS_OS_NAME" = "osx" ]; then
   brew install cmake \
      boost \
      go \
      ccache

   if [ "$RUN_TYPE" = "coverage" ]; then
      brew install lcov
      sudo gem install coveralls-lcov
   fi
fi

pip3 install --user dataclasses-json Jinja2 importlib_resources pluginbase
