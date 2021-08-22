//
// Created by romak on 25.05.2021.
//

#include "UAV.h"

void UAV::setType(UAV::hubType t)
{
	type = t;

	if (type == smallHub)
	{
		meanV = 20 * 5000;
	}

	else if (type == mediumHub)
	{
		meanV = 100 * 5000;
	}

	else
	{
		meanV = 500 * 5000;
	}
}

double UAV::getMeanV() const
{
	return this->meanV;
}
