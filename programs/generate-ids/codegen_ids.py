#!/usr/bin/env python3

import jinja2

import argparse
import hashlib
import json
import os
import sys

def do_hash(name, prefix_value):
    h = int(hashlib.sha256(name.encode("utf-8")).hexdigest(), 16)
    h_top_28 = h >> (256 - 28)
    return hex(h_top_28 | (prefix_value << 28))

# Thunk ID is 8 followed by top 28 bits of H("thunk_id::"+name)
def get_thunk_id(name):
    return do_hash("thunk_id::"+name, 8)

# System call ID is 9 followed by top 28 bits of H("system_call_id::"+name)
def get_system_call_id(name):
    return do_hash("system_call_id::"+name, 9)

def main(argv):

    argparser = argparse.ArgumentParser(description="Generate thunk / system call ID's")

    argparser.add_argument("namefile", metavar="FILE", type=str, help="JSON file containing names of system calls and thunks")
    argparser.add_argument("-o", "--output", metavar="DIR", default=".", type=str, help="Output directory")
    argparser.add_argument("-t", "--template", metavar="DIR", default="templates", type=str, help="Template directory")
    argparser.add_argument("-c", "--check", action="store_true", help="Template directory")

    args = argparser.parse_args(argv)

    with open(args.namefile, "r") as f:
        o = json.load(f)
    system_calls = []
    for name in o["system_calls"]:
        system_calls.append({"name" : name, "system_call_id" : get_system_call_id(name)})
    thunks = []
    for name in o["thunks"]:
        thunks.append({"name" : name, "thunk_id" : get_thunk_id(name)})

    env = jinja2.Environment(
            loader=jinja2.FileSystemLoader(args.template),
            keep_trailing_newline=True,
        )

    ctx = {"system_calls" : system_calls,
           "thunks" : thunks,
          }

    template_names = [
        "thunk_ids.hpp.j2",
        "system_call_ids.hpp.j2",
        ]

    result_files = {}

    for template_name in template_names:
        j2_template = env.get_template(template_name)
        out_filename = os.path.splitext(template_name)[0]
        result_files[out_filename] = j2_template.render(ctx)

    for out_filename in sorted(result_files.keys()):
        fn = os.path.join(args.output, out_filename)
        if os.path.exists(fn):
            with open(fn, "r") as f:
                current_file_contents = f.read()
            if current_file_contents == result_files[out_filename]:
                continue
            if args.check:
                raise RuntimeError("codegen_ids --check failed, files do not match")
        else:
            if args.check:
                raise RuntimeError("codegen_ids --check failed, file does not exist")
        with open(fn, "w") as f:
            f.write(result_files[out_filename])

    return

if __name__ == "__main__":
    main(sys.argv[1:])
