package service

import (
	"fmt"
	"io/ioutil"
	"math"
	"orderManager/internal/domain"
	"strings"
)


func getHubs() (hubs[]domain.Hub, err error){
	data, err := ioutil.ReadFile("configs/validated_hubs.txt")
	if err != nil {
		return nil, err
	}

	dataLines := strings.Split(string(data), "\n")

	for i := 0; i < len(dataLines); i++ {

		if dataLines[i] != "" {
			var (lat, lon float64
				t, index int )
			_, err := fmt.Sscanf(dataLines[i], "%d %f %f %d", &t, &lon, &lat, &index)
			if err != nil {
				return nil, err
			}
			newHub := domain.Hub{Id: index, Latitude: lat * math.Pi / 180, Longitude: lon * math.Pi / 180, TypeHub: t}
			hubs = append(hubs, newHub)
		}
	}


	return hubs, nil
}

func getR(lat1, lat2, lon1, lon2 float64) float64{
	dln := lon2 - lon1
	ch := math.Pow(math.Cos(lat2)*math.Sin(dln), 2) +
		math.Pow(math.Cos(lat1)*math.Sin(lat2)-math.Sin(lat1)*math.Cos(lat2)*math.Cos(dln), 2)
	zn := math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1) * math.Cos(lat2)*math.Cos(dln)

	return math.Atan2(ch, zn) * 6372.795
}