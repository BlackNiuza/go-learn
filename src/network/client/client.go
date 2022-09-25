package main

import (
	"fmt"
	"net"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:51000")
	if err !=nil {
		fmt.Println("Error dailing", err.Error())
		return
	}
	conn.Write([]byte("hello, I'm client"))
}

