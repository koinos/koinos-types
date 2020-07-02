#!/usr/bin/env python3

import hashlib
import json

from collections import OrderedDict

names = [
   "prints",
   "verify_block_header",
   "apply_block",
   "apply_transaction",
   "apply_reserved_operation",
   "apply_upload_contract_operation",
   "apply_execute_contract_operation",
   "apply_set_system_call_operation",
   "db_put_object",
   "db_get_object",
   "db_get_next_object",
   "db_get_prev_object",
   "get_contract_args_size",
   "get_contract_args",
   "exit_contract",
   ]

def read_nonempty_lines(filename):
    result = []
    with open(filename, "r") as f:
        for line in f:
            line = line.strip()
            if line != "":
                result.append(line)
    return result

def do_hash(name, prefix_value):
    h = int(hashlib.sha256(name.encode("utf-8")).hexdigest(), 16)
    h_top_28 = h >> (256 - 28)
    return hex(h_top_28 | (prefix_value << 28))

# Thunk ID is 8 followed by top 28 bits of H("thunk_id::"+name)
def get_thunk_id(name):
    return do_hash("thunk_id::"+name, 8)

# Thunk ID is 9 followed by top 28 bits of H("system_call_id::"+name)
def get_system_call_id(name):
    return do_hash("system_call_id::"+name, 9)

def go():
    with open("json/names.json", "r") as f:
        names_data = json.load(f)
    thunk_names = names_data["thunk_names"]
    system_call_names = names_data["system_call_names"]

    result = OrderedDict()
    result["thunks"] = []
    result["system_calls"] = []

    for name in thunk_names:
        result["thunks"].append(OrderedDict([["name", name], ["thunk_id", get_thunk_id(name)]]))
    for name in system_call_names:
        result["system_calls"].append(OrderedDict([["name", name], ["system_call_id", get_system_call_id(name)]]))

    with open("json/ids.json", "w") as f:
        json.dump(result, f, indent=1)

if __name__ == "__main__":
    go()
