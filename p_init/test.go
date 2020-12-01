package main

import (
	"ext/p_init/lib"
	_ "ext/p_init/lib"
	"fmt"
)

func init() {
	fmt.Println("test.init")
}

func test() {
	fmt.Println(lib.Sum(1, 2, 3))
}
