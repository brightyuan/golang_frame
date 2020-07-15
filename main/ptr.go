package main

import "fmt"

/*
指针使用
*/

func main() {
	var a = 10
	var ip *int //指针声明

	ip = &a /*指针变量被赋值变量a的地址,对指针ip的修改也就是对a的修改*/

	fmt.Printf("a=%d; &a=%d; ip=%d; p2=%d\n", a, &a, ip, *ip)

	*ip = 2 //修改ip <=> 修改a
	fmt.Printf("a=%d; &a=%d; ip=%d; p2=%d\n", a, &a, ip, *ip)

	a = 11 //修改a <=> 修改ip
	fmt.Printf("a=%d; &a=%d; ip=%d; p2=%d\n", a, &a, ip, *ip)

}
