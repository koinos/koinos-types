#!/usr/bin/python3

import os
import shutil
import subprocess

go_cmd = shutil.which('go') + ' run test.go'

p = subprocess.Popen([go_cmd], shell=True)
p.wait()
