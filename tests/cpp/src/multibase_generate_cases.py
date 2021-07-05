#!/usr/bin/env python3

import base58
import base64
import hashlib
import multibase
import random

# pip install py-multibase

def get_test_strings():
    test_strings = [n*b"a" for n in range(10)]
    test_strings += [n*b"\0" for n in range(1, 10)]
    test_strings += [b"this is a test"]
    test_strings += [b"multibase encoding"]
    test_strings += [hashlib.sha256(b"").digest()]
    return test_strings

def mbencode(base, s):
    #
    # We use base58 because py-multibase is known to fail cases with leading zeros
    #
    # See:
    # - https://github.com/multiformats/py-multibase/issues/11
    # - https://github.com/multiformats/multibase/issues/34
    #
    if base == "base58btc":
        return b"z"+base58.b58encode(s)
    if base == "base64pad":
        return b"M"+base64.b64encode(s)
    if base == "base64urlpad":
        return b"U"+base64.urlsafe_b64encode(s)
    return multibase.encode(base, s)

def build_test_case(s):
    #
    # We use base64 from Python stdlib because base64pad is not supported
    result = [
       mbencode("base16", s),
       mbencode("base58btc", s),
       mbencode("base64", s),
       mbencode("base64pad", s),
       mbencode("base64url", s),
       mbencode("base64urlpad", s),
       ]
    result = [e.decode("ascii") for e in result]
    return result

def main():
    first = True
    for s in get_test_strings():
        if first:
            first = False
        else:
            print(",")
        encodings = build_test_case(s)
        print("{" + ", ".join(['"{}"'.format(e) for e in encodings]) + "}", end="")
    print("")

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
