package main


/**
竞态检测——检测代码在并发环境下可能出现的问题
当多线程并发运行的程序竞争访问和修改同一块资源时，会发生竞态问题
go run -race racedetect.go
 */

import (
	"fmt"
	"sync/atomic"
	"time"
)

var seq int64

func GenID() int64 {
	atomic.AddInt64(&seq,1)    //直接return此行，解决竞态检测
	return seq
}

func main() {
	for i := 0; i < 10; i++ {
		go GenID()
	}
	time.Sleep(time.Second*1)
	fmt.Println(GenID())
}

