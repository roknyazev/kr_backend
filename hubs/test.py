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

r = requests.get("http://roknyazev.engineer/state")
data = r.content
print(len(data))
i = struct.unpack('i', data)
print(i)
