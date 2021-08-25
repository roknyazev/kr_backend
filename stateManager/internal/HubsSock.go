package internal

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"net"
	"stateManager/internal/handlers"
	"strconv"
)

func Ff(h *handlers.Handler){

	var c[] net.Conn
	for i := 0; i < 212; i++ {
		conn, _ := net.Dial("tcp", "127.0.0.1"+":"+strconv.Itoa(20000+i))
		c = append(c, conn)
	}
	for  {
		//time.Sleep(20 * time.Millisecond)
		if h.States.TimeBase + 10  < handlers.Timestamp() {
			h.CurrState.Hubs = nil
			h.CurrState.Time = handlers.Timestamp()
			for i := 0; i < 212; i++ {
				var dd handlers.HubReceive
				c[i].Write([]byte("1"))

				byteSlice := make([]byte, 4)
				n := 0
				//message := bufio.NewReader(c[i])
				message := bufio.NewReaderSize(c[i], 228000)
				for n == 0 {
					_, _ = message.Read(byteSlice)
					n = int(binary.LittleEndian.Uint32(byteSlice[:]))
				}
				if n == 1 {
					continue
				}
				//fmt.Printf("received n: %d\n", n)
				byteTrack := make([]byte, n)
				_, _ = message.Read(byteTrack)
				json.Unmarshal(byteTrack, &dd)

				h.CurrState.Hubs = append(h.CurrState.Hubs, dd)
			}

			h.States.TimeBase = handlers.Timestamp()
			h.States.MassStates = append(h.States.MassStates, h.CurrState)
			//fmt.Printf("len h.States.MassStates: %d\n", len(h.States.MassStates))
			//if len(h.CurrState.Hubs) != 0 {
				//fmt.Printf("len h.States.ss: %d\n", len(h.CurrState.Hubs[0].DronesList))
				//fmt.Printf("count hubs: %d\n", len(h.CurrState.Hubs))
			//}
			//fmt.Printf("len TimeBase: %d\n", h.States.TimeBase)
		}
	}
}