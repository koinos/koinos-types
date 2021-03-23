#!/usr/bin/env python3

# One-time migration of test_data.json for #153

import collections
import json

import base58

def encode_varint(x):
    result = []
    while True:
        result.append(x & 0x7F)
        if x <= 0x7F:
            break
        result[-1] |= 0x80
        x = x >> 7
    return bytes(result)

def convert_hook(pairs):
    d = collections.OrderedDict(pairs)
    if set(d.keys()) != set(["hash", "digest"]):
        return d
    id_bytes = encode_varint(d["hash"])
    digest_bytes = base58.b58decode(d["digest"][1:])
    # digest_len_bytes = encode_varint(len(digest_bytes))
    # digest_len_bytes seems to already be included in the data
    result = b"z"+base58.b58encode(id_bytes + digest_bytes)
    if d["digest"] == "zX1ECcx53wEief":
        print("")
        print(d)
        print("id_bytes:", id_bytes.hex())
        # print("digest_len_bytes:", digest_len_bytes.hex())
        print("digest_bytes:", digest_bytes.hex())
        print("concat:", (id_bytes + digest_bytes).hex())
        print("result:", result)
    return result.decode("ascii")

def main():
    with open("json/test_data.json", "r") as f:
        test_data = json.load(f, object_pairs_hook=convert_hook)
    with open("json/test_data.json", "w") as f:
        json.dump(test_data, f, indent=3)

if __name__ == "__main__":
    main()
