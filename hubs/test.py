import json

from sortedcontainers import SortedDict
from collections import defaultdict
import requests
import struct
from typing import Dict, List

from multiprocessing import Process, Queue
import time
from math import *
import pickle
import hashlib

import random

#r = requests.get("http://roknyazev.engineer/state_binary")
#data = r.content
#print(len(data))
#i = struct.unpack('iifffifff', data)
#print(i)

indexes = [1, 2, 3, 4, 5]
print(indexes)

i1 = random.choice(indexes)
print(indexes[0:i1] + indexes[i1 + 1:])
i2 = random.choice(indexes[0:i1] + indexes[i1 + 1:])

print(time.time())