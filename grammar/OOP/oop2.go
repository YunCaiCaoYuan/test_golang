package OOP

import "fmt"

//A是接口
//base -> A(内嵌base)
//base -> B(内嵌base)
//Scene (接口）

type Base2 struct {
	Id   int
	Name string
}

func (*Base2) BasePrint() string {
	return "Base"
}

type Object interface {
	Print() string
}

type A2 struct {
	Base
	AField int
}

func (*A2) Print() string {
	return "A"
}

type B2 struct {
	Base
	BField int
}

func (*B2) Print() string {
	return "B"
}

type Scene2 struct {
	Field1 int
	Field2 string
}

func (*Scene2) EnterScene(arg Object) {
	// 动态绑定
	fmt.Println("EnterScene2:", arg.Print())
}
