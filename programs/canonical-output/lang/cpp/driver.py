#!/usr/bin/python3

import subprocess

command = "./canonical-output-cpp"
p = subprocess.Popen([command])
p.wait()
