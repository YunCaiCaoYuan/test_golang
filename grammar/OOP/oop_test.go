package OOP

import "testing"

// 多态:使用interface
func Test_oop(t *testing.T) {
	obj1 := new(A)

	scene := new(Scene)
	scene.EnterScene(obj1)

	obj2 := new(B)
	scene.EnterScene(obj2)
}

// 多态:使用接口
func Test_oop2(t *testing.T) {
	obj1 := new(A2)

	scene := new(Scene2)
	scene.EnterScene(obj1)

	obj2 := new(B2)
	scene.EnterScene(obj2)
}
