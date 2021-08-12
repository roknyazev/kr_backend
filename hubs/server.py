from http.server import BaseHTTPRequestHandler, HTTPServer
from sortedcontainers import SortedDict
from collections import defaultdict
import json

orders_to_pick_up = []
orders_to_drop_off = []
this_hub_id = None
supply = SortedDict()
base_port = 18000


class ScheduleSlot:
    def __init__(self, order):
        self.time = order.dep_time
        self.dirs = defaultdict(list)
        self.dirs[order.dst_id].append(order)

    def add_order(self, order):
        self.dirs[order.dst_id].append(order)


class Product:
    def __init__(self, post_data):
        self.d = json.loads(post_data)
        self.step = self.d['step'] + 1
        self.oid = self.d['id']
        self.route = self.d['route']
        self.first_pos = self.d['fPos']
        self.last_pos = self.d['lPos']
        self.weight = self.d['weight']
        self.dep_time = self.route[self.step]["Dep_time"]
        self.dst_id = None
        if self.dep_time != 0:
            self.dst_id = self.route[self.step + 1]["HubID"]


class S(BaseHTTPRequestHandler):
    def __init__(self, request, client_address, server):
        super().__init__(request, client_address, server)

    def _set_response(self):
        self.send_response(200)
        self.send_header('Content-type', 'text/html')
        self.end_headers()

    def do_POST(self):
        content_length = int(self.headers['Content-Length'])
        post_data = self.rfile.read(content_length)
        new_order = Product(post_data)
        if new_order.step == 0:
            orders_to_pick_up.append(new_order)
            self._set_response()
            self.wfile.write("Order received".encode('utf-8'))
        if new_order.dep_time != 0:
            try:
                supply[new_order.dep_time].add_order(new_order)
            except KeyError:
                supply[new_order.dep_time] = ScheduleSlot(new_order)
            self._set_response()
            self.wfile.write("Order received".encode('utf-8'))
        else:
            orders_to_drop_off.append(new_order)
            self._set_response()
            self.wfile.write("Order finished".encode('utf-8'))


def run(hub_identifier, address, port, server_class=HTTPServer, handler_class=S):
    global this_hub_id
    this_hub_id = hub_identifier
    server_address = (address, port)
    httpd = server_class(server_address, handler_class)
    print('Starting hub ' + str(this_hub_id) + ' server: ' + str(server_address) + '\n')
    try:
        httpd.serve_forever()
    except KeyboardInterrupt:
        pass
    httpd.server_close()
    print('Stopping hub ' + str(this_hub_id) + ' server: ' + str(server_address) + '\n')
