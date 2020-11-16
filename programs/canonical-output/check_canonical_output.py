#!/usr/bin/env python3

import argparse
import functools
import json
import os
import shutil
import struct
import subprocess
import sys

def check_binary(files, test_data, canon):
   ok = True
   for test in test_data:
      data = {}

      # Read next object for all targets
      for target, f in files.items():
         fbytes = f.read(4)
         if fbytes == '':
            print("Target " + target + " unexpected EOF on test " + test["type"])
            return False

         data_size = struct.unpack('>i', fbytes)[0]

         data[target] = list(f.read(data_size))
         if data[target] == ['']:
            print("Target " + target + " unexpected EOF on test " + test["type"])
            return False

      canon_binary = data[canon]

      # Compare targets to canon
      for target, binary in data.items():
         if target == canon:
            continue

         if len(binary) != len(canon_binary) or not functools.reduce(lambda x, y : x and y, map(lambda p, q : p == q, binary, canon_binary), True):
            ok = False
            print("Target %s does not match canonical serialization on test %s." % (target, test["type"]))
            print("   %s: %s" % (target, str(binary)))
            print("   %s: %s" % (canon, str(canon_binary)))

   for target, f in files.items():
      pos = f.tell()
      f.seek(0, 2)
      if pos != f.tell():
         ok = False
         print("Target %s has %d unexpected bytes at end of file" % (target, f.tell() - pos))

   return ok

def ordered(obj):
    if isinstance(obj, dict):
        return sorted((k, ordered(v)) for k, v in obj.items())
    if isinstance(obj, list):
        return sorted(ordered(x) for x in obj)
    else:
        return obj

def check_json(files, test_data, canon):
   ok = True
   data = {}

   print("Checking json")

   # Load json
   for target, f in files.items():
      try:
         data[target] = json.loads(f.read())
      except JSONDecodeError as err:
         ok = False
         print("Target %s error when loading json: %s" % (target, str(err)))

   # Compare targets to canon
   test_num = 0
   for test in test_data:
      try:
         canon_ordered = ordered(data[canon][test_num])

         for target, arr in data.items():
            if target == canon:
               continue

            if test_num >= len(arr):
               ok = False
               print("Target %s does not have json for test %s" % (target, test["type"]))

            try:
               if ordered(arr[test_num]) != canon_ordered:
                  ok = False
                  print("Target %s does not match json for test %s" % (target, test["type"]))
                  print("   %s: %s" % (target, str(arr[test_num])))
                  print("   %s: %s" % (canon, str(data[canon][test_num])))
            except TypeError:
               ok = False
               print("Target %s does not match json for test %s" % (target, test["type"]))
               print("   %s: %s" % (target, str(arr[test_num])))
               print("   %s: %s" % (canon, str(data[canon][test_num])))
      except TypeError:
         print("Error parsing canon test %s" % test["type"])
         print("canon: %s" % str(data[canon][test_num]))

         for target, arr in data.items():
            if target == canon or test_num >= len(arr):
               continue

            print("%s: %s" % (target, str(arr[test_num])))

      test_num += 1

def main(argv):
   argparser = argparse.ArgumentParser(description="Check Canonical Output")

   argparser.add_argument("-l", "--lang-dir", metavar="DIR", default="", type=str, help="Directory containing canonical output language targets")
   argparser.add_argument("-t", "--test-data", metavar="FILE", default="", type=str, help="File containing json test data")
   argparser.add_argument("-c", "--canon", default="cpp", type=str, help="Target language to consider canon")
   args = argparser.parse_args(argv)

   if args.lang_dir == "":
      sys.exit("Required: lang-lir")

   if args.test_data == "":
      sys.exit("Required: test-file")

   binary_files = {}
   json_files = {}

   python_bin = shutil.which("python3")

   # Run canonical outputs binaries and store output filenames
   for dir_name, sub_dirs, file_list in os.walk(args.lang_dir):
      if dir_name != args.lang_dir and not "CMakeFiles" in dir_name:
         dir_name = os.path.abspath(dir_name)
         target = os.path.split(dir_name)[1]
         print("Running canonical output for %s... " % target, end='')
         p = subprocess.Popen([python_bin + "./driver.py"], cwd=dir_name, shell=True)
         p.wait()

         binary_files[target] = open(os.path.join(dir_name, 'types.bin'), "rb")
         json_files[target] = open(os.path.join(dir_name, 'types.json'), "r")
         print("Done!")

   if not args.canon in binary_files:
      sys.exit("Canon language not found")

   test_data = None
   with open(args.test_data) as json_file:
      test_data = json.load(json_file)

   ok = check_binary(binary_files, test_data, args.canon)
   ok = check_json(json_files, test_data, args.canon) and ok

   return 0 if ok else 1

if __name__ == "__main__":
    result = main(sys.argv[1:])
    sys.exit(result)
