//
// Created by romak on 25.05.2021.
//

#include "Router.h"

std::vector<std::string> split(const std::string &s, char delim)
{
	std::stringstream ss(s);
	std::string item;
	std::vector<std::string> elems;
	while (std::getline(ss, item, delim))
	{
		if (item != "")
			elems.push_back(item);
	}
	return elems;
}

void fill_hub_list(ABalancer::HubList *list)
{
	Hub::hubType sh = Hub::smallHub;
	Hub::hubType mh = Hub::mediumHub;
	Hub::hubType lh = Hub::largeHub;


	int i = 0;
	int type;

	std::ifstream file("../../hubs.txt");
	std::string str;
	std::vector<std::string> data;
	while(getline(file,str))
	{
		data = split(str, ' ');
		type = std::stoi(data[0]);

		if (type == 0)
			list->push_back(new Hub(sh, std::stod(data[1]), std::stod(data[2]), 0));
		if (type == 1)
			list->push_back(new Hub(mh, std::stod(data[1]), std::stod(data[2]), 0));
		if (type == 2)
			list->push_back(new Hub(lh, std::stod(data[1]), std::stod(data[2]), 0));
		i++;
	}
}

Router::~Router()
{

}

Router::Router() :
	priceModel(nullptr, nullptr, &distModel, 0, 0),
	timeModel(nullptr, nullptr, nullptr, &distModel)

{
	fill_hub_list(&list);
	balancer = new GreedyBalancer(list, &distModel, &priceModel, &timeModel);
	balancer->paintGraph("../../graph.eps");
}

struct productData
{
    size_t t;
	double weight;
	int first_hub;
	int last_hub;
};

bool read_message();

[[noreturn]] void Router::start()
{
	int client_sock, listener;
	struct sockaddr_in saddr{};
	int n = sizeof(size_t) + sizeof(double) + sizeof(int) + sizeof(int);
	char buf_res[n];
    char *buf_send = (char *)calloc(list.size() * 16 + 4, sizeof(char));
	std::cout << "Buf size: " << list.size() * 16 + 4 << " bytes" << std::endl;
	productData data{};
	Product product;
	Product::ProductPath path;
	ssize_t err;

	memset(buf_res, 0, n);
	size_t bytes_read;
	listener = socket(AF_INET, SOCK_STREAM, 0);
	if (listener < 0)
	{
		perror("Socket");
		exit(1);
	}
	saddr.sin_family = AF_INET;
	saddr.sin_port = htons(12345);
	saddr.sin_addr.s_addr = htonl(2130706433);
	while (bind(listener, (struct sockaddr *)&saddr, sizeof(saddr)) < 0)
	{
		perror("Bind");
        sleep(1);
	}
	listen(listener, 1);

	std::cout << "Server started!" << std::endl << std::endl;
	while (true)
	{
        client_sock = accept(listener, nullptr, nullptr); //ожидание соединения
        std::cout << "Client socket: " << client_sock << std::endl;
		if(client_sock < 0)
		{
			perror("Accept");
            continue;
		}
		std::cout << "Connection created" << std::endl << std::endl;
		while (true)
		{
			bytes_read = recv(client_sock, buf_res, n, 0);
			if (bytes_read <= 0)
				break;
			memcpy(&data, buf_res, n);
			product.weight = (int)data.weight;
			product.dep = list[data.first_hub];
			product.dst = list[data.last_hub];
			product.dst_time_fhub = data.t;
			//std::cout << data.t;

			balancer->setProductPath(product);
			path = product.getPath();
			int it_buf = 4;
			//std::cout << "Route" << std::endl;
            for (auto & it : path) {
                memcpy(&buf_send[it_buf], &(it.hub->hub_id), 4);
                std::cout << "Hub" << ": " << it.hub->hub_id << " " << it.dest_time << "  " << it.dep_time << std::endl;
                it_buf += 4;
                memcpy(&buf_send[it_buf], &(it.dest_time), 4);
                it_buf += 4;
                memcpy(&buf_send[it_buf], &(it.dep_time), 4);
                it_buf += 4;
            }
            memcpy(&buf_send[0], &(it_buf), 4);
            err = send(client_sock, buf_send, it_buf, MSG_NOSIGNAL);
			if (err == -1)
                break;
			std::cout << "Sent: " << err << " bytes" << std::endl << std::endl;
		}
		close(client_sock);
	}
}
