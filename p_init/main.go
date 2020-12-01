package main

import "fmt"

func init()  {
	fmt.Println("main.init.2")

}
func main() {
	test()
}

func init()  {
	fmt.Println("main.init.1")
}
