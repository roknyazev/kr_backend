from math import *
import time


class Uav:
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
        self.lat_final = lat2
        self.lon_final = lon2

    def dist_calc(self):
        delta_lam = self.lam2 - self.lam1
        tmp1 = sqrt(pow((cos(self.phi2) * sin(delta_lam)), 2) + pow((cos(self.phi1)*sin(self.phi2) - sin(self.phi1) * cos(self.phi2) * cos(delta_lam)), 2))
        tmp2 = sin(self.phi1) * sin(self.phi2) + cos(self.phi1) * cos(self.phi2) * cos(delta_lam)
        result = atan2(tmp1, tmp2) * 6371
        return result


if __name__ == '__main__':
    uavList = []
    for i in range(1):
        uavList.append(Uav(10, 10, 13, 7, 15000))

    t = time.time()
    iters = 0
    while time.time() - t < 1:
        cur_time = time.time()
        i = 0
        for uav in uavList:
            part = cur_time - uav.t0
            print(uav.lat_final - uav.lat_vel * part)
            print(uav.lon_final - uav.lon_vel * part)
            if part > uav.t:
                uavList.pop(i)
                continue
            i += 1
        iters += 1
    print(iters)
