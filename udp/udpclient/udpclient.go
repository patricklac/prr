package main

import (
	"io"
	"log"
	"net"
	"os"
)

// debut, OMIT
const srvAddr = "127.0.0.1:6000"

func main() {
	conn, err := net.Dial("udp", srvAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go func() {
		mustCopy(os.Stdout, conn)
	}()
	mustCopy(conn, os.Stdin) // CTRL-D pour sortir
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

// fin, OMIT
