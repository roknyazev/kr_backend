package handlers

import (
	"time"
)

func Timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}


type (
	Drone struct {
		UID			int64	`json:"uid" binding:"required"`
		T 			int32 	`json:"type" binding:"required"`
		Lat 		float32 `json:"lat" binding:"required"`
		Lon 		float32	`json:"lon" binding:"required"`
		Az      	float32 `json:"az" binding:"required"`
		OrderCount	int32	`json:"count" binding:"required"`
		//Route 	Edge 	`json:"route" binding:"required"`
		// Orders []Order `json:"orders" binding:"required"`
	}
	Order struct {
		Id 		int32	`json:"id" binding:"required"`
		Weight 	float32	`json:"weight" binding:"required"`
		Route	[]Edge	`json:"route" binding:"required"`
	}

	Edge struct {
		//
	}

 	HubReceive struct{
		Id         int32   `json:"id" binding:"required"`
		DronesList []Drone `json:"drones" binding:"required"`
	}

	Compress struct{
		ID int32 `json:"drone" binding:"required"`
		Info Drone `json:"info" binding:"required"`
	}
)

type (
	state struct {
		Time 	int64
		Hubs 	[]HubReceive
	}
	massState struct {
		TimeBase 	int64
		MassStates 	[]state
	}
)
