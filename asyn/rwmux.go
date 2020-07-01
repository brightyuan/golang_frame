package main

import (
	"fmt"
	"sync"
	"time"
)

/**
读写互斥锁（sync.RWMutex）——在读比写多的环境下比互斥锁更高效
同时读不阻塞
 */

var(
	rw_count      int
	rw_countGuard sync.RWMutex
)

func GetRwCount() int  {
	rw_countGuard.RLock()
	defer rw_countGuard.RUnlock()
	return rw_count
}

func setRwCount(c int) {
	rw_countGuard.Lock()
	time.Sleep(time.Second*1)
	rw_count =c
	rw_countGuard.Unlock()
}

func main() {
	setRwCount(10)
	fmt.Println(GetRwCount())
}


