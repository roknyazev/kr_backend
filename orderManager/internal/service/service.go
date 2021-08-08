package service

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"net"
	"orderManager/internal/config"
	"orderManager/internal/domain"
	"orderManager/internal/repository"
	"os"
	_"strconv"
	"time"
)

type Services struct {
	rep *repository.Repository
	Hubs[]domain.Hub
	AddrRouter string
}

func (s *Services) GetNearHub(alPos, lnPos float64) (int, error){
	idHub := -1
	minL := math.Inf(1)
	for i := 0; i < len(s.Hubs); i++ {
		dR := getR(alPos, s.Hubs[i].Latitude, lnPos, s.Hubs[i].Longitude)
		if dR < minL{
			minL 	= dR
			idHub 	= s.Hubs[i].Id
		}
	}

	if  minL > 1 {
		return 0, errors.New("not found")
	}else{
		return idHub, nil
	}
}

func (s *Services) CreateNewOrder(order domain.Order) error{
	_, err := s.rep.CreateNewOrder(order)
	return err
}

func (s *Services) GetTrack(weight float64 , fHub, lHub int32, conn *net.Conn) (string, error){
	for *conn==nil {
		fmt.Println("ненаход, кладмен мудак сука" + s.AddrRouter)
		*conn, _ = net.Dial("tcp", s.AddrRouter)
		if *conn != nil {
			_ = (*conn).SetDeadline(time.Now().Add(time.Hour))
		}
		time.Sleep(500 * time.Millisecond)
	}

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, weight)
	err = binary.Write(buf, binary.LittleEndian, fHub)
	err = binary.Write(buf, binary.LittleEndian, lHub)

	data := make([]byte, 16)
	if err := binary.Read(buf, binary.BigEndian, &data); err != nil {
		return "", err
	}
	n := 0
	for n == 0 {
		n, err = (*conn).Write(data)
		//fmt.Println("n: " + strconv.Itoa(n))
		//fmt.Println(err)
		//n, err = fmt.Fprintf(*conn, string(data[:]))
		if n > 0 && err==nil{
			break
		}

		*conn, _ = net.Dial("tcp", s.AddrRouter)
		time.Sleep(50 * time.Millisecond)
	}


	n = 0
	byteSlice := make([]byte, 4)
	message:= bufio.NewReader(*conn)
	_, _ = message.Read(byteSlice)
	n = int(binary.LittleEndian.Uint32(byteSlice[:]))
	byteTrack := make([]byte, n - 4)
	_, _ = message.Read(byteTrack)

	//for i:=0; i < n - 4; i+=4{
	//	fmt.Println(int(binary.LittleEndian.Uint32(byteTrack[i:i+4])))
	//}

	return "", nil
}


func NewService(r *repository.Repository, c *config.Config) *Services{
	hubs, err := getHubs()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	ss := c.Router.Host + ":" + c.Router.Port
	return &Services{r, hubs, ss}
}

