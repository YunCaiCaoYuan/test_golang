package OOP

import "fmt"

//A是接口
//base -> A(内嵌base)
//base -> B(内嵌base)
//Scene (传参数）

type Base struct {
	Id   int
	Name string
}

func (*Base) BasePrint() string {
	return "Base"
}

type A struct {
	Base
	AField int
}

func (*A) Print() string {
	return "A"
}

type B struct {
	Base
	BField int
}

func (*B) Print() string {
	return "B"
}

type Scene struct {
	Field1 int
	Field2 string
}

func (*Scene) EnterScene(arg interface{}) {
	switch i := arg.(type) {
	case *A:
		fmt.Println("EnterScene:", i.Print())
	case *B:
		fmt.Println("EnterScene:", i.Print())
	}
}
