package chann

import (
	"fmt"
	"time"
)

/*
同步执行
*/
func worker(done chan bool) {
	fmt.Print("working......")
	time.Sleep(time.Second * 1)
	fmt.Print("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	flg := <-done
	fmt.Println("\n", flg)
}
