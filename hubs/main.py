from server import *
import time
import threading
import requests
import multiprocessing


def drop_off(order):
    pass


def pick_up(order):
    pass


def deliver_between_hubs(slot):
    # {'orders': [Order, Order, Order, ...], 'departure_hub_id': id, 'destination_hub_id': id}
    for next_hub, products in slot.dirs.items():
        d = {"orders": [products], 'departure_hub_id': this_hub_id, 'destination_hub_id': next_hub}
        r = requests.post("http://192.168.1.78:8000/api/orders", data=d)


def work():
    while True:
        time.sleep(0.01)
        t = time.time()
        for order in orders_to_pick_up:
            pick_up(order)
        for order in orders_to_drop_off:
            drop_off(order)
        if len(supply) == 0:
            continue
        closest_slot_dep_time = supply.peekitem(0)[0]
        if t >= closest_slot_dep_time:
            deliver_between_hubs(supply.popitem(0)[1])


def run_proc(port):
    thread = threading.Thread(target=work)
    thread.start()
    run(hub_identifier=(port - base_port), address="localhost", port=port)


if __name__ == '__main__':
    processes = []
    for i in range(212):
        processes.append(multiprocessing.Process(target=run_proc, args=(base_port + i,)))
        processes[i].start()
    processes[0].join()
