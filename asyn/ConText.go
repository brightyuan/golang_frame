package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, "hello")
	go worker(ctx, "work")
	go worker(ctx, "bright")

	time.Sleep(time.Second * 2) //先让子线程运行，然后ctx cancel 子线程全部退出，最后main退出
	fmt.Println("stop the gorutine")
	cancel()
	time.Sleep(time.Second * 6)

}

func worker(ctx context.Context, str string) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("sub worker get stop channel")
				return
			default:
				fmt.Println(str, "working")
				time.Sleep(time.Second * 3)
			}
		}
	}()
}
