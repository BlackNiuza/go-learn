package main

import (
	"fmt"
	"net"
)

func doServer(conn net.Conn) {
	for {
		buff:=make([]byte, 512)
		len, err := conn.Read(buff)
		if err!=nil {
			fmt.Println("Error reading", err.Error())
			return 
		}
		fmt.Printf("Received data: %v", string(buff[:len]))
	}
}

func main() {
	fmt.Println("Starting server ...")
	listener, err := net.Listen("tcp", "localhost:51000")
	if err!=nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err!=nil {
			fmt.Println("Error accept", err.Error())
			return
		}
		go doServer(conn)
	}
}
