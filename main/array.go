package main

import "fmt"

//初始化
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
var balance2 = [5]float32{1000.0, 2.0, 3.4, 7.0}

var n = [10]int{12, 7, 9}

func main() {
	for i := 0; i < len(balance2); i++ {
		fmt.Println(balance2[i])
	}
	fmt.Println(balance)
	fmt.Println(n)

	fmt.Println(n[0] / n[1])
	fmt.Println(float32(n[0]) / float32(n[1])) //变量类型转换
}
