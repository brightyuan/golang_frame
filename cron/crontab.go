package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	go cronJobPerSeconds()
	go cronJobPerMin()
	go cronJobPerDay()
	select {}
}

//每s执行一次
func cronJobPerSeconds() {
	ct := cron.New(cron.WithSeconds())
	ct.AddFunc("*/1 * * * * *", func() {
		fmt.Println("hello cron")
	})
	ct.Start()
	//使main进程阻塞 select{} 和  <-make(chan int)
	select {}
	//<-make(chan int)
}

//每分钟执行一次
func cronJobPerMin() {
	c := cron.New()
	spec := "0 */1 * * * *"
	c.AddFunc(spec, func() {
		fmt.Println("execute per second")
	})
	c.Start()
	select {}
}

// 每天上午9点到12点的第2和第10分钟执行
func cronJobPerDay() {
	spec := "2,10 9-12 * * *"
	c := cron.New()
	c.AddFunc(spec, myFunc)
	c.Start()
	select {}
}

func myFunc() {
	fmt.Println("executed！")
}
