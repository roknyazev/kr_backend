package handler

import (
	_"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"orderManager/internal/domain"
)

func (h *Handler) CreateNewOrder(c *gin.Context){

	var newOrder domain.Order

	if err := c.BindJSON(&newOrder); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	fHub, err := h.srv.GetNearHub(newOrder.FirstLat * math.Pi / 180 ,newOrder.FirstLon * math.Pi / 180)
	lHub, err := h.srv.GetNearHub(newOrder.LastLat * math.Pi / 180,newOrder.LastLon * math.Pi / 180)

	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	_, err = h.srv.GetTrack(newOrder.Weight, int32(fHub), int32(lHub), h.con)
	if err != nil {
		return
	}

	//fmt.Print("ordeeeeeeeeeeeeeeeer")
	c.JSON(http.StatusOK, newOrder)

}

func (h *Handler) GetAllOrders(c *gin.Context) {
	return
}

func (h *Handler) GetAllHubs(c *gin.Context) {
	return
}


