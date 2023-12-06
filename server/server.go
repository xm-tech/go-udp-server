package main

import (
	"log"
	"net"
)

func main() {
	go serve()

	select {}
}

func serve() {
	lis, err := net.ListenUDP("udp", &net.UDPAddr{Port: 3333, IP: net.ParseIP("127.0.0.1")})
	if err != nil {
		log.Printf("listen err: %+v\n", err)
		return
	}
	defer lis.Close()

	log.Println("server started")

	for {
		var data [1024]byte
		n, addr, err := lis.ReadFromUDP(data[:])
		if err != nil {
			log.Println("read err,err:", err)
			continue
		}
		_, _ = lis.WriteToUDP(data[:n], addr)

		log.Printf("got:%+v,remote:%+v, bn:%+v", string(data[:n]), addr, n)
	}
}
