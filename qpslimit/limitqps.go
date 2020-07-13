package main

import (
	"fmt"
	"time"
)

/**
服务器限流：通过令牌桶实现限流
1.通过创建带有buffer 的 chan，实现令牌桶，向令牌桶中添加fillToken()
2.通过takeToken() 消费令牌
*/

func main() {
	capacity := 100
	tokenBucket := make(chan struct{}, capacity)

	//添加限流chan
	interVal := time.Microsecond * 100
	go fillToken(tokenBucket, interVal)

	//模拟用户调用
	for {
		go takeToken(tokenBucket)
		time.Sleep(time.Microsecond * 50)
	}

	select {} //main阻塞
}

func fillToken(tokenBucket chan struct{}, fillInterval time.Duration) {
	ticker := time.NewTicker(fillInterval)
	for {
		select {
		case <-ticker.C:
			select {
			case tokenBucket <- struct{}{}:
			default:
			}
			fmt.Println("current token cnt:", len(tokenBucket), time.Now())
		}
	}
}

func takeToken(tokenBucket chan struct{}) bool {
	var takeResult bool
	select {
	case <-tokenBucket:
		takeResult = true
	default:
		takeResult = false
	}
	fmt.Println("take token res:", takeResult)
	return takeResult
}
