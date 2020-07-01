package main

/**
结构体+方法
定义结构体union，实现handle1，handle2两个方法
 */
import "fmt"

type union2 struct {
	i int
	t string
}

func (u *union2) HandleMethod() {
	fmt.Println("echo inter1")

}

func (u *union2) HandleMethod2() {
	fmt.Println("echo inter2")

}

func main() {
	u := union2{i: 12, t: "hello"}
	u.HandleMethod()
	u.HandleMethod2()
}
