package chann

import (
	"fmt"
	"time"
)

func sender(c chan string)  {
	t:= time.NewTicker(time.Second*1)
	for ; ;  {
		c<- "send a msg"
		<-t.C
	}
}


func main() {
	c1 := make(chan string)
	stop := make(chan bool)
	go sender(c1)
	go func() {
		time.Sleep(time.Second*3)
		stop<-true
	}()

	for ; ;  {
		select {      //只执行最先返回的消息
		case msg := <-c1:
			fmt.Println("recv ", msg)
		case  <-stop:
			fmt.Println("time up!")
			return
		}
	}

}
