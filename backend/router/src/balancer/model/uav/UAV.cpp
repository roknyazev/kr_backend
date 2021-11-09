//
// Created by romak on 25.05.2021.
//

#include "UAV.h"

void UAV::setType(UAV::hubType t)
{
	type = t;

	if (type == smallHub)
	{
		meanV = 2000;
	}

	else if (type == mediumHub)
	{
		meanV = 3500;
	}

	else
	{
		meanV = 5000;
	}
}

double UAV::getMeanV() const
{
	return this->meanV;
}
