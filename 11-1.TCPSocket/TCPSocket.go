package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	address := net.TCPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8000,
		Zone: "",
	}
	listener, err := net.ListenTCP("tcp4", &address)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Other Address: ", conn.RemoteAddr())
		go echo(conn)
	}
}
func echo(conn *net.TCPConn) {
	tick := time.Tick(5 * time.Second)
	for now := range tick {
		n, err := conn.Write([]byte(now.String()))
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		fmt.Printf("send %d bytes to %s\n", n, conn.RemoteAddr())
	}
}
