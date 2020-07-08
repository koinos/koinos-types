#!/usr/bin/env python3

import hashlib
import sys

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

def main():
    if len(sys.argv) != 2:
        print("Usage: generate_ids.py <thunk_name>")
        return
    
    name = sys.argv[1]
    thunk_id = get_thunk_id(name)
    system_call_id = get_system_call_id(name)

    print(f"Thunk:          {name}")
    print(f"Thunk ID:       {thunk_id}")
    print(f"System Call ID: {system_call_id}")

    return

if __name__ == "__main__":
    main()
