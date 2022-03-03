package struct_test

import (
	"fmt"
	"testing"
)

type a struct {
	a1 int
}

func (*a) say() {
	fmt.Println("a saying")
}

type b struct {
	a
	b1 int
}

func newObj() *b {
	return &b{}
}

/*
继承
一个结构体嵌到另一个结构体，称作组合
匿名和组合的区别
如果一个struct嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现继承
如果一个struct嵌套了另一个【有名】的结构体，那么这个模式叫做组合
如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现多重继承
*/

func Test_Obj(t *testing.T) {
	obj := newObj()
	obj.say()
}
