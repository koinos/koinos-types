#!/usr/bin/python3

import os
import shutil
import subprocess

gopath = os.popen('go env GOPATH').read()[:-1] + ":" + os.path.join(os.getcwd(), "../../../../generated/golang")
gocache = os.popen('go env GOCACHE').read()
go_cmd = shutil.which('go') + ' run test.go'

p = subprocess.Popen([go_cmd], env={"GOPATH": gopath, "GOCACHE": gocache, "HOME": os.environ["HOME"]}, shell=True)
p.wait()
