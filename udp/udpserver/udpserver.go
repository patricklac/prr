package main

import (
	"log"
	"net"
)

const srvAddr = "127.0.0.1:6000"

// debut, OMIT
func main() {
	addr, _ := net.ResolveUDPAddr("udp", srvAddr)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, cliAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		s := string(buf[0:n-1]) + " from " + cliAddr.String() + "\n"
		if _, err := conn.WriteTo([]byte(s), cliAddr); err != nil {
			log.Fatal(err)
		}
	}
}

// fin, OMIT
