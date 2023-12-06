package main

import (
	"log"
	"net"
)

func main() {
	run()
}

func run() {
	raddr := net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 3333}
	so, err := net.DialUDP("udp", nil, &raddr)
	if err != nil {
		log.Println("dialUDP err:", err)
		return
	}
	defer so.Close()

	udata := []byte("hello")
	log.Printf("send:%+v", string(udata))
	_, err = so.Write(udata)
	if err != nil {
		log.Println("send data fail, err:", err)
		return
	}

	rdata := make([]byte, 1024)
	bn, remote, err := so.ReadFromUDP(rdata)
	if err != nil {
		log.Println("readFromUDP err:", err)
		return
	}
	log.Printf("got:%+v,remote:%+v,bn:%+v", string(rdata), remote, bn)
}
