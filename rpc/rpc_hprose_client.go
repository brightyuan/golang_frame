package main

/**
安装hprose
go get "github.com/hprose/hprose-golang/rpc"
*/

import (
	"fmt"
	"github.com/hprose/hprose-golang/rpc"
)

func main() {
	type clientStub struct {
		Hello3 func(string) string //注意方法名大写
	}
	client := rpc.NewTCPClient("tcp://127.0.0.1:8288/")
	defer client.Close()
	var stub clientStub
	client.UseService(&stub)
	rep := stub.Hello3("world")
	fmt.Println(rep)
}
