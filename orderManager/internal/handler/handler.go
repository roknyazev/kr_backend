package handler

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net"
	_"orderManager/internal/domain"
	"orderManager/internal/service"
	"sync"
	"time"
)

type Handler struct {
	mu	sync.Mutex
	srv *service.Services
	con        *net.Conn
	idOrder    int32
	MassOrders sync.Map
}

func NewHandler(s *service.Services) *Handler {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil && conn != nil{
		err = conn.SetDeadline(time.Now().Add(time.Second))
		fmt.Println(err)
	}
	return &Handler{srv: s, con: &conn, idOrder: 0}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "хуй соси"})
	})

	ordG := router.Group("/orders")
	{
		ordG.POST("/new", h.CreateNewOrder)
		ordG.GET("/all", h.GetAllOrders)
		ordG.POST("/update", h.UpdOrder)
		ordG.GET("/getOrderInfo", h.GetOrder)
		ordG.GET("/getOrders", h.GetOrders)
	}
	hubG := router.Group("/hubs")
	{
		hubG.GET("/all", h.GetAllHubs)
	}
	return router
}




