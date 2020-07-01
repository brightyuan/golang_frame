package main

import "fmt"

/*
指针使用
*/

func main() {

	var a = 10
	var ip *int    //指针声明
	ip = &a       /*指针变量存储地址*/

	var p2 *string

	fmt.Println(ip)
	fmt.Println(*ip)
	fmt.Println(a)
	fmt.Println(p2)

}
