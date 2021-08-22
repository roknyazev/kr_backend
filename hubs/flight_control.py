from math import *
import time
import requests
from server import *
import socket
import json
import struct
import numpy as np
import multiprocessing


drone_list = []


class OneFlight:
    def __init__(self, lat1, lon1, lat2, lon2, vel):
        self.phi1 = lat1 * pi / 180  # rad
        self.lam1 = lon1 * pi / 180  # rad
        self.phi2 = lat2 * pi / 180  # rad
        self.lam2 = lon2 * pi / 180  # rad
        self.vel = vel  # km/h
        self.distance = self.dist_calc()  # km
        self.t = (self.distance / self.vel) * 3600  # sec
        self.t0 = time.time()
        self.lat_vel = (self.lam2 - self.lam1) / self.t
        self.lon_vel = (self.phi2 - self.phi1) / self.t
        self.lat_final = lat2 * pi / 180
        self.lon_final = lon2 * pi / 180
        self.lat_start = lat1 * pi / 180
        self.lon_start = lon1 * pi / 180

    def dist_calc(self):
        delta_lam = self.lam2 - self.lam1
        tmp1 = sqrt(pow((cos(self.phi2) * sin(delta_lam)), 2) + pow((cos(self.phi1)*sin(self.phi2) - sin(self.phi1) * cos(self.phi2) * cos(delta_lam)), 2))
        tmp2 = sin(self.phi1) * sin(self.phi2) + cos(self.phi1) * cos(self.phi2) * cos(delta_lam)
        result = atan2(tmp1, tmp2) * 6371
        return result


class Drone:
    def __init__(self, path, drone_type, order_list):
        self.step = 0
        self.path = path
        self.drone_type = drone_type
        print("DRONE TYPE:", drone_type)
        if drone_type <= 0:
            self.velocity = 20
        elif drone_type == 1:
            self.velocity = 200
        elif drone_type == 2:
            self.velocity = 500
        self.order_list = order_list
        self.start_time = None
        self.flight_list = []
        for j in range(len(path) - 1):
            self.flight_list.append(OneFlight(path[j][1], path[j][0], path[j + 1][1], path[j + 1][0], self.velocity))
        self.start(time.time())

        self.cur_lat = 0
        self.cur_lon = 0

    def start(self, start_time):
        self.start_time = start_time

    def get_position(self, curr_time):
        one_flight = self.flight_list[self.step]
        time_part = curr_time - one_flight.t0
        if time_part > one_flight.t:
            self.step += 1
            self.start_time = curr_time
            try:
                self.flight_list[self.step].t0 = self.start_time
            except IndexError:
                return -1, -1, -1
            return one_flight.lon_final, one_flight.lat_final, 0

        north_dir = np.array([1, 1])
        cur_dir = np.array([self.cur_lon, self.cur_lat])
        self.cur_lat = one_flight.lat_start + one_flight.lat_vel * time_part
        self.cur_lon = one_flight.lon_start + one_flight.lon_vel * time_part
        next_dir = np.array([self.cur_lon, self.cur_lat])
        az = np.dot(cur_dir - next_dir, north_dir) / np.linalg.norm(cur_dir - next_dir)
        return self.cur_lon, self.cur_lat, az + 0.5 * pi


def sim(drone_queue: multiprocessing.Queue, supply_queue: multiprocessing.Queue):
    print('Starting flight control server ' + str(this["id"]) + '\n')
    hub_id = this["id"]
    sock = socket.socket()
    while True:
        try:
            sock.bind(('localhost', flight_control_base_port + hub_id))
            break
        except OSError:
            print("Address localhost:" + str(flight_control_base_port + hub_id) + " already in use")
            time.sleep(1)
            continue
    sock.listen(1)
    conn, _ = sock.accept()
    try:
        while True:
            while not drone_queue.empty():
                drone_list.append(drone_queue.get())

            try:
                conn.recv(1)
            except ConnectionResetError:
                continue

            # time.sleep(0.01)
            coord_lst = []
            i = 0
            while i < len(drone_list):
                drone = drone_list[i]
                coord = drone.get_position(time.time())
                if coord == (-1, -1, -1):
                    if drone_list[i].drone_type == -1:
                        supply_queue.put(drone_list[i].order_list[0])
                    elif drone_list[i].drone_type == -2:
                        drone_list.pop(i)
                        continue
                    else:
                        lst = drone_list[i].order_list
                        next_hub = lst[0].dst_id
                        for order in lst:
                            requests.post('http://localhost:' + str(order_receiver_base_port + next_hub), json=order.d)
                    drone_list.pop(i)
                    continue
                coord_lst.append({"type": drone_list[i].drone_type, "lat": (coord[1] / pi) * 180, "lon": (coord[0] / pi) * 180, "az": (coord[2] / pi) * 180})
                i += 1
            if len(coord_lst) != 0:
                data = {"id": this["id"], "drones": coord_lst}
                send = json.dumps(data).encode("ascii")
                try:
                    conn.send(struct.pack("i" + "{}s".format(len(send)), len(send), bytes(send)))
                except BrokenPipeError:
                    conn, _ = sock.accept()
                # print(data)
            else:
                try:
                    conn.send(struct.pack("i", 1))
                except BrokenPipeError:
                    conn, _ = sock.accept()
    except KeyboardInterrupt:
        print('Stopping flight control server ' + str(this["id"]) + '\n')
        conn.close()
