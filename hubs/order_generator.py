import requests
import random
import time


min_lat = 53
max_lat = 58

min_lon = 88
max_lon = 98

t = time.time()

coord_list = []
validated_hubs = open("validated_hubs.txt")
for hub_info in validated_hubs:
    hub_type, longitude, latitude, _ = hub_info.strip().split()
    coord_list.append({"lon": float(longitude), "lat": float(latitude)})
n = len(coord_list)
indexes = list(range(n))

tmp = [[], []]

while True:
    time.sleep(.1)
    i1 = random.choice(indexes)
    i2 = random.choice(indexes[0:i1] + indexes[i1 + 1:])
    print(i1, i2)

    fs = coord_list[i1]
    ls = coord_list[i2]
    first_lat = fs["lat"] + random.uniform(-0.1, 0.1)
    first_lon = fs["lon"] + random.uniform(-0.1, 0.1)

    last_lat = ls["lat"] + random.uniform(-0.1, 0.1)
    last_lon = ls["lon"] + random.uniform(-0.1, 0.1)

    weight = 1
    send = {
        "weight": weight,
        "first_lat": first_lat,
        "first_lon": first_lon,
        "last_lat": last_lat,
        "last_lon": last_lon
        }

    r = requests.post('http://localhost:8080/orders/new', json=send)
    #print()
    ##print(i)
    #print(time.time() - t)
    #print(first_lat, first_lon, last_lat, last_lon)
    #print(r.text)
