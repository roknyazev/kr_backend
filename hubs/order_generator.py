import requests
import random
import time
import json


min_lat = 53
max_lat = 58

min_lon = 88
max_lon = 98

t = time.time()

for i in range(100):
    print(i)

    first_lat = random.uniform(min_lat, max_lat)
    first_lon = random.uniform(min_lon, max_lon)
    last_lat = random.uniform(min_lat, max_lat)
    last_lon = random.uniform(min_lon, max_lon)
    r = requests.post('http://localhost:8080/orders/new', json={
                                                                "weight": 1,
                                                                "first_lat": first_lat,
                                                                "first_lon": first_lon,
                                                                "last_lat": last_lat,
                                                                "last_lon": last_lon
                                                                })
print(time.time() - t)
print(r)
