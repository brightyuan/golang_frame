package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*等待组（sync.Wait Group）——保证在并发环境中完成指定数量的任*/

func main() {
	//声明等待组
	var wg sync.WaitGroup
	var urls = []string{
		"https://www.baidu.com",
		"https://www.sina.com.cn/",
	}
	for _, url := range urls {
		//开始新任务，等待组+1
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, error := http.Get(url)
			fmt.Println(url, error)
		}(url)
	}
	//等待所有任务完成
	wg.Wait()
	fmt.Println("end")
}
