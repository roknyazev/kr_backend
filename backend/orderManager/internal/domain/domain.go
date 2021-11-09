package domain

type Order struct {
	Id       	int32 		`json:"-"`
	Weight		float64		`json:"weight" binding:"required"`
	FirstLat 	float64 	`json:"first_lat" binding:"required"`
	FirstLon 	float64 	`json:"first_lon" binding:"required"`
	LastLat 	float64 	`json:"last_lat" binding:"required"`
	LastLon 	float64 	`json:"last_lon" binding:"required"`
	DroneID		int64 		`json:"-"`
	Route		[]HubTime	`json:"-"`
}

type OrderFront struct {
	Id       	int32 		`json:"-"`
	Weight		float64		`json:"weight" binding:"required"`
	FirstLat 	float64 	`json:"first_lat" binding:"required"`
	FirstLon 	float64 	`json:"first_lon" binding:"required"`
	LastLat 	float64 	`json:"last_lat" binding:"required"`
	LastLon 	float64 	`json:"last_lon" binding:"required"`
	DroneID		int64 		`json:"droneID"`
	Route		[]HubTime	`json:"route"`
}

type SlotOrder struct {
	Id int32 `json:"id" binding:"required"`
	DroneID int64 `json:"droneID" binding:"required"`
}

type Hub struct {
	Id        	int
	Latitude  	float64
	Longitude 	float64
	TypeHub		int
}

type HubTime struct {
	HubId	int32	`json:"HubID" binding:"required"`
	DepTime	int64	`json:"Dep_time" binding:"required"`
	DstTime	int64	`json:"Dst_time" binding:"required"`
}

type OrderSend struct {
	Step	int16		`json:"step" binging:"required"`
	Id		int32		`json:"id" binding:"required"`
	Weight	float64    	`json:"weight" binding:"required"`
	LPos	[2]float64 	`json:"lPos" binding:"required"`
	FPos	[2]float64 	`json:"fPos" binding:"required"`
	Route	[]HubTime  	`json:"route" binding:"required"`
}
type HubTimeForFront struct {
	HubId	int32	`json:"HubID" binding:"required"`
	DepTime	int64	`json:"Dep_time" binding:"required"`
	DstTime	int64	`json:"Dst_time" binding:"required"`
	LonPos	float64	`json:"lon_pos"`
	LanPos 	float64	`json:"lan_pos"`
}

type OrderSendForFront struct {
	Id       	int32 		`json:"order_id" binding:"required"`
	Weight		float64		`json:"weight" binding:"required"`
	FirstLat 	float64 	`json:"first_lat" binding:"required"`
	FirstLon 	float64 	`json:"first_lon" binding:"required"`
	LastLat 	float64 	`json:"last_lat" binding:"required"`
	LastLon 	float64 	`json:"last_lon" binding:"required"`
	DroneID		int64 		`json:"droneID"`
	Route		[]HubTimeForFront	`json:"route"`
}

