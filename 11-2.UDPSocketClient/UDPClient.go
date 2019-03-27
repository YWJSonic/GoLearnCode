package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Fprintf(os.Stderr, "address: %s host:port", os.Args[0])
	// 	os.Exit(1)
	// }
	for {
		service := "127.0.0.1:1200" //os.Args[1]
		udpAddr, err := net.ResolveUDPAddr("udp4", service)
		checkError(err)
		conn, err := net.DialUDP("udp", nil, udpAddr)
		checkError(err)
		_, err = conn.Write([]byte("1234567890"))
		checkError(err)
		var buf [512]byte
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(buf[0:n])
		// os.Exit(0)
		time.Sleep(1 * time.Second)
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
}
