#!/bin/bash

set -e
set -x

if [ -z "$KOINOS_PATH" ]
then
    KOINOS_PATH=../../../koinos
fi

if [ ! -e "$KOINOS_PATH" ]
then
    echo "KOINOS_PATH does not exist"
    exit 1
fi

python3 ./generate_ids.py
mkdir -p out
python3 "$KOINOS_PATH/programs/build_helpers/koinos_build_helpers/buildj2.py" --template-dir templates --json-dir json --output-dir out

cp out/system_call_ids.hpp ../../libraries/pack/src/system_call_ids.hpp
cp out/thunk_ids.hpp ../../libraries/pack/src/thunk_ids.hpp
