package main

import (
	"fmt"
	"io/ioutil"
)
import (
	"strings"
	"time"
)

func main() {
	s := "i am an \n string,just guess\n\t is a tab\n"
	s1 := "i am an" + "hello "
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(len("hello"))
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.ToLower("HELLO"))
	//读取文件错误
	readfileError()
	//加go，非阻塞，并发执行
	go sleeDo()
	fmt.Println("not wait sleep 2s")
	time.Sleep(time.Second * 3)

	c := make(chan string)
	c <- "slow channel tag"
	msg := <-c
	println(msg)
}

func readfileError() {
	file, err := ioutil.ReadFile("foo.txt")
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println("file:  %s", file)
	name, role := "richard Jupp", "drum"
	err1 := fmt.Errorf("this %v %v quit", role, name)
	if err1 != nil {
		fmt.Println(err1)
	}
}

func sleeDo() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleep 2s done")
}
