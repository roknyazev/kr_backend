import threading
from flight_control import *


def drop_off(order: Order, drone_queue):
    path = [hub_list[this["id"]]["position"], order.last_pos, hub_list[this["id"]]["position"]]
    drone_type = -2
    drone_queue.put(Drone(path, drone_type, [order]))
    # drone_list.append(Drone(path, drone_type, [order]))


def pick_up(order: Order, drone_queue):
    path = [hub_list[this["id"]]["position"], order.first_pos, hub_list[this["id"]]["position"]]
    drone_type = -1
    drone_queue.put(Drone(path, drone_type, [order]))
    # drone_list.append(Drone(path, drone_type, [order]))


def deliver_between_hubs(slot: ScheduleSlot, drone_queue):
    # {'orders': [Order, Order, Order, ...], 'departure_hub_id': id, 'destination_hub_id': id}
    for next_hub, orders in slot.dirs.items():
        path = [hub_list[this["id"]]["position"], hub_list[next_hub]["position"]]
        drone_type = min(hub_list[this["id"]]["type"], hub_list[next_hub]["type"])
        drone_queue.put(Drone(path, drone_type, orders))


def work(drone_queue, supply_queue):
    while True:
        time.sleep(0.1)  # (0.01)
        t = time.time()
        for j in range(len(this["orders to pick up"])):
            pick_up(this["orders to pick up"][j], drone_queue)
            this["orders to pick up"].pop(j)
        for j in range(len(this["orders to drop off"])):
            drop_off(this["orders to drop off"][j], drone_queue)
            this["orders to drop off"].pop(j)

        while not supply_queue.empty():
            print("SUPPLY              fshifuisahuifhiohfuosdhiofiosd")
            order = supply_queue.get()
            try:
                this["supply"][order.dep_time].add_order(order)
            except KeyError:
                this["supply"][order.dep_time] = ScheduleSlot(order)
        if len(this["supply"]) == 0:
            continue
        closest_slot_dep_time = this["supply"].peekitem(0)[0]
        if t >= closest_slot_dep_time:
            deliver_between_hubs(this["supply"].popitem(0)[1], drone_queue)


def run_proc(port: int):
    drone_queue = multiprocessing.Queue()
    supply_queue = multiprocessing.Queue()

    thread1 = threading.Thread(target=work, args=(drone_queue, supply_queue))
    thread1.start()
    this["id"] = port - order_receiver_base_port
    process = multiprocessing.Process(target=sim, args=(drone_queue, supply_queue))
    process.start()
    run(ip="localhost", port=port)
    # thread2 = threading.Thread(target=sim)
    # thread2.start()


if __name__ == '__main__':
    processes = []
    validated_hubs = open("validated_hubs.txt")
    for hub_info in validated_hubs:
        hub_type, longitude, latitude, _ = hub_info.strip().split()
        hub_list.append({"type": int(hub_type), "position": (float(longitude), float(latitude))})
    for i in range(len(hub_list)):
        processes.append(multiprocessing.Process(target=run_proc, args=(order_receiver_base_port + i,)))
        processes[i].start()
    try:
        processes[0].join()
    except KeyboardInterrupt:
        print("Stop")
