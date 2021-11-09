//
// Created by romak on 25.05.2021.
//

#include "Hub.h"

Hub::Hub(Hub::hubType type, double x, double y, unsigned int hub_id) : type(type), posX(x * M_PI / 180.), posY(y * M_PI / 180.), lat(x), lon(y), hub_id(hub_id)
{
	if (type == smallHub)
	{
		max_distance = 150;
		max_load_capacity = 100;
		color[0] = 0;
		color[1] = 0;
		color[2] = 255;
		step = 50;
	}

	else if (type == mediumHub)
	{
		max_distance = 500;
		max_load_capacity = 1000;
		color[0] = 0;
		color[1] = 255;
		color[2] = 0;
		step = 100;
	}

	else
	{
		max_distance = 5000;
		max_load_capacity = 10000;
		color[0] = 255;
		color[1] = 0;
		color[2] = 0;
		step = 200;
	}

}

