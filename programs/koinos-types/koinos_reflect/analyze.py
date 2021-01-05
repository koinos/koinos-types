
import argparse
import json
import os
import sys

from . import lexer
from . import parser
from . import schema

def read_files(filenames):
    result = []
    if filenames is None:
        return result
    for fn in filenames:
        with open(fn, "r") as f:
            content = f.read()
            result.append(content)
            if not content.endswith("\n"):
                result.append("\n")
    return "".join(result)

def write_if_different(dest_filename, content):
    if dest_filename == "-":
        sys.stdout.write(content)
        sys.stdout.flush()
        return
    if os.path.exists(dest_filename):
        with open(dest_filename, "r") as f:
            current_content = f.read()
        if current_content == content:
            return

    with open(dest_filename, "w") as f:
        f.write(content)

def main(argv):
    argparser = argparse.ArgumentParser(description="Analyze files")

    argparser.add_argument("typefiles",   metavar="FILES", nargs="*", type=str, help="Type file(s)")
    argparser.add_argument("-l", "--lex",    action="store_true", help="Lex input file(s)")
    argparser.add_argument("-p", "--parse",  action="store_true", help="Parse input file(s)")
    argparser.add_argument("-s", "--schema", action="store_true", help="Generate schema")
    argparser.add_argument("-o", "--output", default="-", type=str, help="Output file (default=stdout)")
    args = argparser.parse_args(argv)

    content = read_files(args.typefiles)

    output = []
    if args.lex:
        for tok in lexer.lex(content):
            output.append(tok.to_json(separators=(",", ":")))
            output.append("\n")
    if args.parse:
        output.append(parser.parse(content).to_json(separators=(",", ":")))
        output.append("\n")
    if args.schema:
        parsed_content = parser.parse(content)
        schemanator = schema.Schemanator(parsed_content)
        schemanator.schemanate()
        schema_json = schemanator.schema.to_json(separators=(",", ":"))
        # Make sure the JSON we produce is loadable with from_json()
        schema2 = schema.Schema.from_json(schema_json)
        output.append(schema_json)
        output.append("\n")

    write_if_different(args.output, "".join(output))

if __name__ == "__main__":
    main(sys.argv[1:])
