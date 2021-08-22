from sortedcontainers import SortedDict
from collections import defaultdict
import requests
from typing import Dict, List

from multiprocessing import Process, Queue
import time


def f(q):
    q.put(3)
    q.put(3)
    q.put(3)


if __name__ == '__main__':
    q = Queue()
    p = Process(target=f, args=(q,))
    p.start()
    time.sleep(0.1)
    while not q.empty():
        print(q.get())
    p.join()


