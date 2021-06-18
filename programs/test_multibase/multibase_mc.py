#!/usr/bin/env python3

import base58
import base64
import json
import multibase
import random
import sys

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

def generate_byte(rand):
    ZERO_CHANCE = 25
    if rand.randrange(0, 100) < ZERO_CHANCE:
        return 0
    return rand.randrange(1, 0x100)

def generate_length(rand):
    LONG_CHANCE = 3
    if rand.randrange(0, 100) < LONG_CHANCE:
        return rand.randrange(10, 1000)
    return rand.randrange(3, 10)

def generate_case(i, n):
    # linear shuffle by some arbitrary constants to use as random seed
    rand = random.Random(0xb501*i + 0xfff1*n)
    u = []
    for i in range(generate_length(rand)):
        u.append(generate_byte(rand))
    return bytes(u)

# TODO:  Test all possible 1-char and 2-char elements

def main():
    bases = ["base16", "base58btc", "base64", "base64pad", "base64url", "base64urlpad"]

    i = 0
    runid = 1
    while True:
        c = generate_case(runid, i)
        u = [mbencode(b, c).decode("ascii") for b in bases]
        try:
            sys.stdout.write(json.dumps(u)+"\n")
        except BrokenPipeError as e:
            break
        i += 1

if __name__ == "__main__":
    main()
