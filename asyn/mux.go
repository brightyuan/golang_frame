package main

import (
	"fmt"
	"sync"
	"time"
)

/**
互斥锁（sync.Mutex）——保证同时只有一个goroutine可以访问共享资源
 */
var(
	count int
	countGuard sync.Mutex
)

func GetCount() int  {
	//rw_countGuard.Lock()
	//defer rw_countGuard.Unlock()
	return count
}

func setCount(c int) {
	countGuard.Lock()
	time.Sleep(time.Second*1)
	count=c
	countGuard.Unlock()
}

func main() {
	setCount(10)
	fmt.Println(GetCount())
}

