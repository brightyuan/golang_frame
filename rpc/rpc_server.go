package main

/**
安装hprose
go get "github.com/hprose/hprose-golang/rpc"
*/

import (
	"github.com/hprose/hprose-golang/rpc"
)

func main() {
	server := rpc.NewTCPServer("tcp4://0.0.0.0:8288/")
	server.AddFunction("Hello3", hello3)
	server.Start()
}

func hello3(s string) string {
	return s + " from Tcp server"
}
