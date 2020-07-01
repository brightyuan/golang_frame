package chann

import (
	"fmt"
	"time"
)

func ping1(c chan string) {
	time.Sleep(time.Second *2)
	c <- "ping on channel 1"
}

func ping2(c chan string) {
	time.Sleep(time.Second * 1)
	c <- "ping on channel 2"
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go ping1(c1)
	go ping2(c2)

	select {      //只执行最先返回的消息
	case msg := <-c1:
		fmt.Println("recv ", msg)
	case msg2 := <-c2:
		fmt.Println("recv ", msg2)
	case <-time.After(500*time.Millisecond):          //500ms内未收到消息，采取措施
		fmt.Println("no msg recv ,stop")
	}
}
