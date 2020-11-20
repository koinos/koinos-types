#!/usr/bin/python3

import subprocess
import os

cmd = os.getcwd()
for d in ['Debug', 'Release', 'RelWithDebInfo', 'MinSizeRel']:
   if os.path.isdir(os.path.join(os.getcwd(), d)):
      cmd = os.path.join(cmd, d)
      break

command = cmd + "/canonical-output-cpp"
p = subprocess.Popen([command])
p.wait()
