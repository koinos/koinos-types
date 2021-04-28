#!/bin/bash

set -e
set -x

rm -Rf build
mkdir -p build
python3 -m koinos_reflect.analyze ../../types/basetypes.def ../../types/block_store_rpc.hpp ../../types/block_store.hpp ../../types/broadcast.hpp ../../types/chain_rpc.hpp ../../types/chain.hpp ../../types/common.hpp ../../types/mempool_rpc.hpp ../../types/protocol.hpp -s -o build/block.schema
python3 -m koinos_codegen.codegen --target-path lang --target typescript -o build -p koinos_protocol build/block.schema

# cp tester.py build
# cd build
# python3 tester.py
