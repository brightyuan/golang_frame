package chann

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	t := time.NewTicker(1 * time.Second)   //间隔1s 触发一次计算器
	for {
		c <- "ping"
		<-t.C
	}

}

func main() {
	message := make(chan string)
	go pinger(message)
	//for ; ;  {
	//	msg := <- message
	//	fmt.Println(msg)
	//}

	for i := 0; i < 5; i++ {
		msg := <-message
		fmt.Println(msg)
	}

}


/*
c : = make(chan string)     //1.创建通道
c<- "hello world"              // 2. 写消息
msg:= <- c             			//3.读消息
*/