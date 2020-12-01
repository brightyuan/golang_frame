package main

import (
	"fmt"
	"net/http"
	"time"
)

//参数数量不定，num[0]
func test(num ...int) {
	fmt.Println(num)
	var sum = 0
	for _, v1 := range num {
		sum = sum + v1
	}
	fmt.Println(sum)
}

func test2() error {
	var n = 1 / 1
	fmt.Println(n)
	return nil
}

func test3() {
	// get the location
	location, _ := time.LoadLocation("GMT")
	// this should give you time in location
	//t := time.Now().In(location).Format("02 Jan 2006 15:04:05 GMT")
	t := time.Now().In(location).Format(http.TimeFormat)
	fmt.Println(t)
}

func main() {
	//test(1,2)
	//test(1,2,5,9)
	//e := test2()
	//panic(e)
	test3()
}
