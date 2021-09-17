package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"orderManager/internal/domain"
	"sort"
	"strconv"
	"time"
)

func (h *Handler) CreateNewOrder(c *gin.Context){
	h.mu.Lock()
	h.idOrder++

	var newOrder domain.Order

	if err := c.BindJSON(&newOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		h.mu.Unlock()
		return
	}

	fHub, DFirstMile, err := h.srv.GetNearHub(newOrder.FirstLat * math.Pi / 180, newOrder.FirstLon * math.Pi / 180)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		h.mu.Unlock()
		return
	}
	lHub, _, err := h.srv.GetNearHub(newOrder.LastLat * math.Pi / 180,newOrder.LastLon * math.Pi / 180)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		h.mu.Unlock()
		return
	}

	timeCopter := (DFirstMile / float64(h.srv.Configs.Router.VCopter)) * 3600 * 2 + float64(h.srv.Configs.Router.TimeWaitOrder)


	now := time.Now()
	OS, err := h.srv.GetTrack(now.Unix() + int64(timeCopter) ,newOrder.Weight, int32(fHub), int32(lHub), h.con)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		fmt.Println(err)
		h.mu.Unlock()
		return
	}

	var ord domain.OrderFront
	ord.Id =  h.idOrder
	ord.Weight = newOrder.Weight
	ord.FirstLat = newOrder.FirstLat
	ord.FirstLon = newOrder.FirstLon
	ord.LastLon = newOrder.LastLon
	ord.LastLat = newOrder.LastLat
	ord.Route = OS.Route
	ord.DroneID = -int64(ord.Route[0].HubId)

	h.MassOrders.Store(ord.Id,&ord)
	//h.MassOrders[ord.Id] = &ord


	OS.Step = -1
	OS.Id = h.idOrder
	OS.FPos[0] = newOrder.FirstLon
	OS.FPos[1] = newOrder.FirstLat
	OS.LPos[0] = newOrder.LastLon
	OS.LPos[1] = newOrder.LastLat
	OS.Weight = newOrder.Weight
	bytesRepresentation, err := json.Marshal(OS) // ERROR
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		fmt.Println(err)
		h.mu.Unlock()
		return
	}
	_, err = http.Post(fmt.Sprintf("http://%s:%d", h.srv.Configs.Hubs.Host ,h.srv.Configs.Hubs.BasePort + fHub),
		"application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		fmt.Println("6")
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		fmt.Println(err)
		h.mu.Unlock()
		return
	}

	var aa domain.OrderSendForFront
	aa.Id = OS.Id
	aa.Weight = OS.Weight
	aa.DroneID = 0
	aa.FirstLat = OS.FPos[1]
	aa.FirstLon = OS.FPos[0]
	aa.LastLat = OS.LPos[1]
	aa.LastLon = OS.LPos[0]

	for i:=0 ; i < len(OS.Route); i++{
		var ee domain.HubTimeForFront
		ee.HubId = OS.Route[i].HubId
		ee.DstTime = OS.Route[i].DstTime
		ee.DepTime = OS.Route[i].DepTime
		for j:=0; j < len(h.srv.Hubs);j++{
			if ee.HubId == int32(h.srv.Hubs[j].Id){
				ee.LanPos = h.srv.Hubs[j].Latitude
				ee.LonPos = h.srv.Hubs[j].Longitude
				break
			}
		}
		aa.Route = append(aa.Route, ee)
	}
	c.JSON(http.StatusOK, aa)
	h.mu.Unlock()
}

func (h *Handler) GetAllOrders(c *gin.Context) {
	return
}

func (h *Handler) GetAllHubs(c *gin.Context) {
	send := h.srv.Hubs
	c.JSON(http.StatusOK, send)
}

func (h *Handler) UpdOrder(c *gin.Context) {
	var updOrder domain.SlotOrder

	if err := c.BindJSON(&updOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	aa1, _ := h.MassOrders.Load(updOrder.Id)
	var ee1 *domain.OrderFront = aa1.(*domain.OrderFront)
	ee1.DroneID = updOrder.DroneID
	//h.MassOrders[updOrder.Id].DroneID = updOrder.DroneID      // invalid memory address or nil pointer dereference

	c.JSON(http.StatusOK, "ok")
}

func (h *Handler) GetOrder(c *gin.Context) {
	type ww struct {
		Id int32 `json:"id"  binding:"required"`
	}
	var idOrd ww
	tmp, _ := strconv.ParseInt(c.Request.URL.Query()["id"][0], 10, 32)
	idOrd.Id = int32(tmp)

	Orderl, _ := h.MassOrders.Load(idOrd.Id) // FIXME
	//Order = *(h.MassOrders[idOrd.Id])
	if Orderl == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Order does not exist")
		return
	}
	var aa domain.OrderSendForFront
	aa.Id = Orderl.(*domain.OrderFront).Id
	aa.Weight = Orderl.(*domain.OrderFront).Weight
	aa.DroneID = Orderl.(*domain.OrderFront).DroneID
	aa.FirstLat = Orderl.(*domain.OrderFront).FirstLat
	aa.FirstLon = Orderl.(*domain.OrderFront).FirstLon
	aa.LastLat = Orderl.(*domain.OrderFront).LastLat
	aa.LastLon = Orderl.(*domain.OrderFront).LastLon

	for i:=0 ; i < len(Orderl.(*domain.OrderFront).Route); i++{
		var ee domain.HubTimeForFront
		ee.HubId = Orderl.(*domain.OrderFront).Route[i].HubId
		ee.DstTime = Orderl.(*domain.OrderFront).Route[i].DstTime
		ee.DepTime = Orderl.(*domain.OrderFront).Route[i].DepTime
		for j:=0; j < len(h.srv.Hubs);j++{
			if ee.HubId == int32(h.srv.Hubs[j].Id){
				ee.LanPos = h.srv.Hubs[j].Latitude
				ee.LonPos = h.srv.Hubs[j].Longitude
				break
			}
		}
		aa.Route = append(aa.Route, ee)
	}

	c.JSON(http.StatusOK, aa)
}

func (h *Handler) GetOrders(c *gin.Context) {
	type ww struct {
		Id int64 `json:"id"  binding:"required"`
	}
	var idOrd ww
	idOrd.Id, _ = strconv.ParseInt(c.Request.URL.Query()["id"][0], 10, 64)

	var orders []int32

	f := func ( key, value interface{} ) bool {
		if value.(*domain.OrderFront).DroneID == idOrd.Id{
			orders = append(orders, key.(int32))
		}
		return true
	}

	h.MassOrders.Range(f)
	//for key, value := range h.MassOrders {
	//	if value.DroneID == idOrd.Id{
	//		orders = append(orders, key)
	//	}
	//}
//
	sort.Slice(orders, func(i, j int) bool { return orders[i] < orders[j] })

	fmt.Println("Orders Send: ")
	fmt.Println(orders)
	c.JSON(http.StatusOK, orders)
}
