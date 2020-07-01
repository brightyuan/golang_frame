package main

import "fmt"

//参数数量不定，num[0]
func test(num ...int) {
	fmt.Println(num)
	var sum = 0
	for _,v1 := range num {
		sum = sum + v1
	}
	fmt.Println(sum)
}

func test2()  error{
	var n = 1 / 0
	fmt.Println(n)
}

func main() {
	test(1,2)
	test(1,2,5,9)
	e := test2()
	panic(e)
}