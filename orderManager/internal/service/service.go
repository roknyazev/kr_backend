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
	"os"
	_"strconv"
	"time"
)

type Services struct {
	Hubs[]domain.Hub
	Configs * config.Config
	AddrRouter string
}

func (s *Services) GetNearHub(alPos, lnPos float64) (int, float64, error){
	idHub := -1
	minL := math.Inf(1)
	for i := 0; i < len(s.Hubs); i++ {
		dR := GetR(alPos, s.Hubs[i].Latitude, lnPos, s.Hubs[i].Longitude)
		if dR < minL{
			minL 	= dR
			idHub 	= s.Hubs[i].Id
		}
	}

	if  minL > float64(s.Configs.Router.MaxDistanceCopter) {
		return 0, 0,errors.New("not found")
	}else{
		return idHub, minL,nil
	}
}

func (s *Services) GetTrack(timeStart int64, weight float64 , fHub, lHub int32, conn *net.Conn) (*domain.OrderSend, error){
	var orderSend domain.OrderSend
	orderSend.Weight = weight

	for *conn==nil {
		fmt.Println("ненаход, кладмен мудак сука" + s.AddrRouter)
		*conn, _ = net.Dial("tcp", s.AddrRouter)
		if *conn != nil {
			_ = (*conn).SetDeadline(time.Now().Add(time.Hour))
		}
		//time.Sleep(500 * time.Millisecond)
	}

	buf := bytes.NewBuffer([]byte{})
	err := binary.Write(buf, binary.LittleEndian, timeStart)
	err = binary.Write(buf, binary.LittleEndian, weight)
	err = binary.Write(buf, binary.LittleEndian, fHub)
	err = binary.Write(buf, binary.LittleEndian, lHub)

	data := make([]byte, 24)
	if err := binary.Read(buf, binary.BigEndian, &data); err != nil {
		return nil, err
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
		//time.Sleep(50 * time.Millisecond)
	}


	n = 0
	byteSlice := make([]byte, 4)
	message:= bufio.NewReader(*conn)

	start := time.Now()
	_, err = message.Read(byteSlice)
	if err!=nil{
		fmt.Println(err)
	}
	duration := time.Since(start)
	fmt.Println(duration.Milliseconds())

	n = int(binary.LittleEndian.Uint32(byteSlice[:]))
	byteTrack := make([]byte, n - 4)
	_, err = message.Read(byteTrack)
	if err!=nil{
		fmt.Println(err)
	}

	countHubTime := (n - 3) / 12
	start = time.Now()
	for i:= 0; i < countHubTime; i++{
		ii := i * 12
		hubId 	:= int32(binary.LittleEndian.Uint32(byteTrack[ii:ii+4]))
		dstTime := int32(binary.LittleEndian.Uint32(byteTrack[ii+4:ii+8]))
		depTime := int32(binary.LittleEndian.Uint32(byteTrack[ii+8:ii+12]))
		hubTime := domain.HubTime{HubId: hubId, DepTime: int64(depTime), DstTime: int64(dstTime)}
		orderSend.Route = append(orderSend.Route, hubTime)
	}
	duration = time.Since(start)
	fmt.Println(duration.Microseconds())
	//for i:=0; i < n - 4; i+=4{
	//	fmt.Println(int(binary.LittleEndian.Uint32(byteTrack[i:i+4])))
	//}

	return &orderSend, nil
}

func NewService( c *config.Config) *Services{
	hubs, err := getHubs()
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}

	ss := c.Router.Host + ":" + c.Router.Port
	return &Services{ hubs,c, ss}
}

