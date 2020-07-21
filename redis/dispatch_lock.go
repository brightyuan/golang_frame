package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

/**
通过redis的setnx实现简单分布式锁设置

*/

func main() {
	client := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:32769",
			Password: "",
			DB:       0,
		})
	fmt.Println(client.Ping())

	//加锁
	resp := client.SetNX("keylock", 1, time.Second*10)
	sus, _ := resp.Result()
	fmt.Println(sus)

	//解锁
	getResp := client.Get("keylock")
	cntVal, _ := getResp.Int64()
	fmt.Println(cntVal)

	resp2 := client.SetNX("keylock", 1, time.Second*10)
	sus2, _ := resp2.Result()
	fmt.Println(sus2)
	fmt.Println(resp2)
}
