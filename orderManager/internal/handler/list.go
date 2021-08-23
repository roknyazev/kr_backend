package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"orderManager/internal/domain"
	"strconv"
	"time"
)

func (h *Handler) CreateNewOrder(c *gin.Context){

	var newOrder domain.Order

	if err := c.BindJSON(&newOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	fHub, DFirstMile, err := h.srv.GetNearHub(newOrder.FirstLat * math.Pi / 180, newOrder.FirstLon * math.Pi / 180)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	lHub, _, err := h.srv.GetNearHub(newOrder.LastLat * math.Pi / 180,newOrder.LastLon * math.Pi / 180)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	timeCopter := (DFirstMile / float64(h.srv.Configs.Router.VCopter)) * 3600 * 2 + float64(h.srv.Configs.Router.TimeWaitOrder)
	fmt.Println("timeCopter " + fmt.Sprint(timeCopter))
	fmt.Println("R " + fmt.Sprint(DFirstMile))

	now := time.Now()
	fmt.Println("текущее время: " + strconv.FormatInt(now.Unix(), 10))
	OS, err := h.srv.GetTrack(now.Unix() + int64(timeCopter) ,newOrder.Weight, int32(fHub), int32(lHub), h.con)
	if err != nil {
		fmt.Println(err)
		return
	}

	(*OS).Step = -1
	(*OS).Id = 0
	(*OS).FPos[0] = newOrder.FirstLon
	(*OS).FPos[1] = newOrder.FirstLat
	(*OS).LPos[0] = newOrder.LastLon
	(*OS).LPos[1] = newOrder.LastLat
	(*OS).Weight = newOrder.Weight
	bytesRepresentation, err := json.Marshal(*OS)
	_, err = http.Post(fmt.Sprintf("http://%s:%d", h.srv.Configs.Hubs.Host ,h.srv.Configs.Hubs.BasePort + fHub),
		"application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, newOrder)
}

func (h *Handler) GetAllOrders(c *gin.Context) {
	return
}

func (h *Handler) GetAllHubs(c *gin.Context) {
	send := h.srv.Hubs
	c.JSON(http.StatusOK, send)
}


