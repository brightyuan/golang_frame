package main

import (
	"net"
	"time"
)

func main() {
	conn, _ := net.Dial("tcp", ":9000")
	//conn, err := net.DialTimeout("tcp", ":8080", 2 * time.Second)
	defer conn.Close()
	conn.Write([]byte("发消息给tcp server"))
	time.Sleep(time.Second * 3)
}
