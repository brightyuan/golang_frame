package main

import (
	"fmt"
	"net"
)

func main() {
	listen, _ := net.Listen("tcp", ":9000")
	for {
		acc, err := listen.Accept()
		if err != nil {
			fmt.Println("error listen", err)
		}
		go handle(acc)
	}
}

func handle(c net.Conn) {
	defer c.Close()
	for {
		fmt.Printf("acc")
	}
}
