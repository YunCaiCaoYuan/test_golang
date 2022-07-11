package OOP

import (
	"fmt"
	"testing"
	"unsafe"
)

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

// 打印 struct{} 大小
func Test_size_struct(t *testing.T) {
	size := unsafe.Sizeof(struct{}{})
	fmt.Println("size:", size) // 0
}
