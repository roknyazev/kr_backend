package handlers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	_ "strconv"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h *Handler) wshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	var SubTime int64 = 0
	var LastTimeSend int64 = 0

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		if msg != nil{
			SubTime = int64(binary.BigEndian.Uint64(msg))
		}

		if SubTime != 0 && LastTimeSend != h.States.MassStates[len(h.States.MassStates) - 1].Time{
			err = conn.WriteMessage(t, msg)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}
}

type Handler struct {
	CurrState state
	States massState
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "хуй соси"})
	})

	router.POST("/", h.setState)

	stateGroup := router.Group("/state")
	{
		stateGroup.GET("/binary", h.GetStateBinary)
		stateGroup.GET("/json", h.GetStateJson)
	}


	router.GET("/ws", func(c *gin.Context) {
		h.wshandler(c.Writer, c.Request)
	})


	return router
}

func (h *Handler) setState(c *gin.Context) {

	var resData HubReceive

	if err := c.BindJSON(&resData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	var check bool = false
	for i := 0; i < len(h.CurrState.Hubs); i++ {
		if h.CurrState.Hubs[i].Id == resData.Id{
			check = true
			break
		}
	}
	if !check {
		h.CurrState.Hubs = append(h.CurrState.Hubs, resData)
	}
	if h.States.TimeBase == 0 {
		h.States.TimeBase = Timestamp()
	}



	if h.States.TimeBase + 1000000000000 < Timestamp(){
		h.CurrState.Time = Timestamp()
		h.States.MassStates = append(h.States.MassStates, h.CurrState)
		h.CurrState.Hubs = nil
		h.CurrState.Time = 0
		h.States.TimeBase = Timestamp()
		// fmt.Printf("len h.States.MassStates: %d\n", len(h.States.MassStates))
		// fmt.Printf("len TimeBase: %d\n", h.States.TimeBase)
	}

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) GetStateBinary(c *gin.Context){
	if len(h.States.MassStates) == 0 {
		c.JSON(http.StatusOK, nil)
		return
	}
	hubs := (h.States.MassStates[len(h.States.MassStates)-1]).Hubs


	var length int32 = 0
	for i := 0; i < len(hubs); i++ {
		length += int32(len(hubs[i].DronesList)*28)
	}
	//fmt.Printf("len length %d \n", length)

	//fmt.Printf("LEN STATES: %d \n", len(h.States.MassStates))
	//fmt.Printf("LEN HUBS: %d \n", len(hubs))

	buf := bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.LittleEndian, length)
	for i := 0; i < len(hubs); i++ {
		drones := hubs[i].DronesList
		for i := 0; i < len(drones); i++ {
			binary.Write(buf, binary.LittleEndian, drones[i].UID) // 8
			binary.Write(buf, binary.LittleEndian, drones[i].T) // 4
			binary.Write(buf, binary.LittleEndian, drones[i].Lon) // 4
			binary.Write(buf, binary.LittleEndian, drones[i].Lat) // 4
			binary.Write(buf, binary.LittleEndian, drones[i].Az) // 4
			binary.Write(buf, binary.LittleEndian, drones[i].OrderCount) // 4
		}
	}
	//fmt.Printf("len buf: %d \n", buf.Len())
	dd := c.Writer
	dd.Write(buf.Bytes())
	dd.WriteHeader(http.StatusOK)
}

func (h *Handler) GetStateJson(c *gin.Context){
	if len(h.States.MassStates) == 0 {
		c.JSON(http.StatusOK, nil)
		return
	}
	hubs := (h.States.MassStates[len(h.States.MassStates)-1]).Hubs
	var send  []Compress
	for i := 0; i < len(hubs); i++ {
		drones := hubs[i].DronesList
		for i := 0; i < len(drones); i++ {
			send = append(send, Compress{int32(i), drones[i]})
		}
	}
	c.JSON(http.StatusOK, send)
}