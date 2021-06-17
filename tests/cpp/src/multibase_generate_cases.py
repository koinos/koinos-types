#!/usr/bin/env python3

import hashlib
import multibase

# pip install py-multibase

def get_test_strings():
    test_strings = [n*b"a" for n in range(10)]
    test_strings += [n*b"\0" for n in range(1, 10)]
    test_strings += [b"this is a test"]
    test_strings += [b"multibase encoding"]
    test_strings += [hashlib.sha256(b"").digest()]
    return test_strings

def build_test_case(s):
    result = [
       multibase.encode("base16", s),
       multibase.encode("base58btc", s),
       multibase.encode("base64", s),
       multibase.encode("base64url", s),
       ]
    result = [e.decode("ascii") for e in result]
    return result

def main():
    for s in get_test_strings():
        encodings = build_test_case(s)
        print("{" + ", ".join(['"{}"'.format(e) for e in encodings]) + "},")

    '''
    with open("test.txt", "r") as f:
        for line in f:
            line = line.strip()
            if line == "":
                continue
            u = line.split(",")
            count = int(u[0], 16)
            if len(u) != count+1:
                raise RuntimeError("bad data file")
            print(create_test_case(u[1:]))
'''

if __name__ == "__main__":
    main()
