package internal

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"stateManager/internal/handlers"
	"strconv"
)

func Ff(h *handlers.Handler){
	var dd handlers.HubReceive
	var c[] net.Conn
	for i := 0; i < 212; i++ {
		conn, _ := net.Dial("tcp", "127.0.0.1"+":"+strconv.Itoa(20000+i))
		c = append(c, conn)
	}
	for  {
		if h.States.TimeBase + 10  < handlers.Timestamp() {
			h.CurrState.Hubs = nil
			h.CurrState.Time = handlers.Timestamp()
			for i := 0; i < 212; i++ {

				c[i].Write([]byte("1"))

				byteSlice := make([]byte, 4)
				n := 0
				message := bufio.NewReader(c[i])
				for n == 0 {
					_, _ = message.Read(byteSlice)
					n = int(binary.LittleEndian.Uint32(byteSlice[:]))
				}
				if n == 1 {
					continue
				}
				byteTrack := make([]byte, n)
				_, _ = message.Read(byteTrack)
				json.Unmarshal(byteTrack, &dd)

				h.CurrState.Hubs = append(h.CurrState.Hubs, dd)
			}
			h.States.TimeBase = handlers.Timestamp()
			h.States.MassStates = append(h.States.MassStates, h.CurrState)
			fmt.Printf("len h.States.MassStates: %d\n", len(h.States.MassStates))
			fmt.Printf("len h.States.ss: %d\n", len(h.CurrState.Hubs))
			fmt.Printf("len TimeBase: %d\n", h.States.TimeBase)
		}
	}
	//or{

	//	if (h.States.TimeBase + 1000 < Timestamp())




	//	var resData hubResive

	//	if err := c.BindJSON(&resData); err != nil {
	//		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	//		return
	//	}

	//	var check bool = false
	//	for i := 0; i < len(h.CurrState.Hubs); i++ {
	//		if h.CurrState.Hubs[i].Id == resData.Id{
	//			check = true
	//			break
	//		}
	//	}
	//	if !check {
	//		h.CurrState.Hubs = append(h.CurrState.Hubs, resData)
	//	}
	//	if h.States.TimeBase == 0 {
	//		h.States.TimeBase = Timestamp()
	//	}

	//	if h.States.TimeBase + 1000 < Timestamp(){
	//		h.CurrState.time = Timestamp()
	//		h.States.MassStates = append(h.States.MassStates, h.CurrState)
	//		h.CurrState.Hubs = nil
	//		h.CurrState.time = 0
	//		h.States.TimeBase = Timestamp()
	//		fmt.Printf("len h.States.MassStates: %d\n", len(h.States.MassStates))
	//		fmt.Printf("len TimeBase: %d\n", h.States.TimeBase)
	//	}


//	}
}