package domain

type Order struct {
	Id       	int 		`json:"-"`
	Weight		float64		`json:"weight" binding:"required"`
	FirstLat 	float64 	`json:"first_lat" binding:"required"`
	FirstLon 	float64 	`json:"first_lon" binding:"required"`
	LastLat 	float64 	`json:"last_lat" binding:"required"`
	LastLon 	float64 	`json:"last_lon" binding:"required"`
}

type Hub struct {
	Id        	int
	Latitude  	float64
	Longitude 	float64
	TypeHub		int
}