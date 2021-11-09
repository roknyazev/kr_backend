from http.server import BaseHTTPRequestHandler, HTTPServer
from sortedcontainers import SortedDict
from collections import defaultdict
import json
import requests


order_receiver_base_port = 18000
flight_control_base_port = 20000
hub_list = []
this = {"id": -1,
        "orders to pick up": list(),
        "orders to drop off": list(),
        "supply": SortedDict()}


class Order:
    def __init__(self, post_data: bytes):
        self.d = json.loads(post_data)
        self.d['step'] += 1
        self.step = self.d['step']
        self.oid = self.d['id']
        self.route = self.d['route']
        self.first_pos = tuple(self.d['fPos'])
        self.last_pos = tuple(self.d['lPos'])
        self.weight = self.d['weight']
        self.dep_time = self.route[self.step]["Dep_time"]
        self.dst_id = None
        if self.dep_time != 0:
            self.dst_id = self.route[self.step + 1]["HubID"]
        self.d['droneID'] = -this['id']
        try:
            requests.post('http://localhost:8080/orders/update', json={"id": self.oid, "droneID": self.d["droneID"]})
        except BaseException:
            print("/orders/update not available")

    def update(self, drone_id):
        self.d['droneID'] = drone_id
        try:
            requests.post('http://localhost:8080/orders/update', json={"id": self.oid, "droneID": self.d["droneID"]})
        except BaseException:
            print("/orders/update not available")


class ScheduleSlot:
    def __init__(self, order: Order):
        self.time = order.dep_time
        self.dirs = defaultdict(list)
        self.dirs[order.dst_id].append(order)

    def add_order(self, order: Order):
        self.dirs[order.dst_id].append(order)


class OrderReceiver(BaseHTTPRequestHandler):
    def __init__(self, request, client_address, server):
        super().__init__(request, client_address, server)

    def _set_response(self):
        self.send_response(200)
        self.send_header('Content-type', 'text/html')
        self.end_headers()

    def do_POST(self):
        content_length = int(self.headers['Content-Length'])
        post_data = self.rfile.read(content_length)
        #print("HUB ", this["id"])
        #print(post_data)
        new_order = Order(post_data)
        if new_order.step == 0:
            this["orders to pick up"].append(new_order)
            self._set_response()
            self.wfile.write("Order received".encode('utf-8'))
        elif new_order.dep_time != 0:
            try:
                this["supply"][new_order.dep_time].add_order(new_order)
            except KeyError:
                this["supply"][new_order.dep_time] = ScheduleSlot(new_order)
            self._set_response()
            self.wfile.write("Order received".encode('utf-8'))
        else:
            this["orders to drop off"].append(new_order)
            self._set_response()
            self.wfile.write("Order finished".encode('utf-8'))

    def log_message(self, format, *args):
        return


def run(ip: str, port: int, server_class=HTTPServer, handler_class=OrderReceiver):

    address = (ip, port)
    httpd = server_class(address, handler_class)
    print('Starting order receiver server ' + str(this["id"]) + ' address: ' + str(address) + '\n')
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        httpd.server_close()
        print('Stopping order receiver server ' + str(this["id"]) + ' address: ' + str(address) + '\n')
