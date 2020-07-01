package main
/**
接口多继承
定义结构体union，继承接口handle1，handle2两个接口
 */
import "fmt"

type union struct {
	i int
	t string
}

type handleI interface {
	HandleMethod()
}

type handleI2 interface {
	HandleMethod2()
}

func (u *union) HandleMethod() {
	fmt.Println("echo inter1")

}

func (u *union) HandleMethod2() {
	fmt.Println("echo inter2")

}

func main() {
	u:= union{i:12,t:"hello"}
	u.HandleMethod()
	//union.HandleMethod2()
	u.HandleMethod2()
}