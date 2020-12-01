package main

/**
结构体+方法
定义结构体union，实现handle1，handle2两个方法
*/
import (
	"fmt"
)

type union2 struct {
	i int
	t string
}

//非指针方式，无法修改结构体内部值
func (u union2) HandleMethod() {
	u.i *= 2
	u.t += " not ptr"
}

//指针方式，可以修改结构体内部值，建议使用！
func (u *union2) HandlePtrMethod() {
	u.i *= 2
	u.t += " ptr"
}

func main() {
	u := union2{
		i: 12,
		t: "hello",
	}
	u.HandleMethod()
	fmt.Println(u)

	u.HandlePtrMethod()
	fmt.Println(u, "结构体值被修改了")
}
