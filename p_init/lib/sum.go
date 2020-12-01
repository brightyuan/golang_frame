package lib

import "fmt"

func init() {
	fmt.Println("lib/sum.init")
}

func Sum(x ...int) int {
	n := 0
	for i := range x {
		n += i
	}
	return n
}
