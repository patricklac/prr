package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"golang.org/x/net/ipv4"
)

// debut, OMIT
const multicastAddr = "224.0.0.1:6000"

func main() {
	addr, err := net.ResolveUDPAddr("udp", multicastAddr)
	if err != nil {
		log.Fatal(err)
	}
	go clientReader(addr)
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(conn, os.Stdin)
}

// milieu, OMIT
func clientReader(addr *net.UDPAddr) {
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	p := ipv4.NewPacketConn(conn)
	if err = p.JoinGroup(nil, addr); err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024)
	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Fatal(err)
		}
		s := bufio.NewScanner(bytes.NewReader(buf[0:n]))
		for s.Scan() {
			fmt.Printf("%s from %v\n", s.Text(), addr)
		}
	}
}

// fin, OMIT
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
