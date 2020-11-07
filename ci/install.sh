#!/bin/bash

if [ "$TRAVIS_OS_NAME" = "linux" ]; then
   sudo apt-get install -y \
      libboost-all-dev \
      python3 \
      python3-pip \
      python3-setuptools \
      valgrind \
      ccache

   eval "$(gimme 1.15.4)"
elif [ "$TRAVIS_OS_NAME" = "osx" ]; then
   brew install cmake \
      boost \
      ccache

   if [ "$RUN_TYPE" = "coverage" ]; then
      brew install lcov
      sudo gem install coveralls-lcov
   fi
fi
eval "$(gimme 1.15.4)"
pip3 install --user dataclasses-json Jinja2 importlib_resources pluginbase
