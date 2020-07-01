package chann

import (
	"fmt"
	"time"
)

func main() {

	//messages := make(chan string)

	//channel buffer 通过make创建，关键字chan指定
	messages := make(chan string, 3)
	fmt.Println(cap(messages))     //返回buffer大小
	fmt.Println(len(messages))   //返回chan中未读消息数量

	go func() { messages <- "ping" }()
	//panic("ok to stop anyway")   //强制停止

	messages <- "go"
	messages <- "go2"
	time.Sleep(time.Second*2)
	close(messages) //关闭通道
	receiver(messages)
	messages <- "go3"  //关闭通道不能再接收消息
	msg := <-messages
	fmt.Println(msg)
}

func receiver(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}
