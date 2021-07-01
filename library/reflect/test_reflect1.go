package main

import "fmt"
import "reflect"

type Student struct {
	Name string
	Age  int
}

func main() {
	//定义一个Student类型的指针变量
	var stu = &Student{Name:"kitty", Age: 20}
	//获取结构体实例的反射类型对象
	typeOfStu := reflect.TypeOf(stu)
	//显示反射类型对象的名称和种类
	fmt.Printf("name: '%v', kind: '%v'\n", typeOfStu.Name(), typeOfStu.Kind())
	//取类型的元素
	typeOfStu = typeOfStu.Elem()
	//显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfStu.Name(), typeOfStu.Kind())
//	name: '', kind: 'ptr'
//	element name: 'Student', element kind: 'struct'
}

/*
//定义一个Enum类型
type Enum int
const (
	Zero Enum = 0
)
type Student struct {
	Name string
	Age  int
}
func main() {
	//定义一个Student类型的变量
	var stu Student
	//获取结构体实例的反射类型对象
	typeOfStu := reflect.TypeOf(stu)
	//显示反射类型对象的名称和种类
	fmt.Println(typeOfStu.Name(), typeOfStu.Kind())
	//获取Zero常量的反射类型对象
	typeOfZero := reflect.TypeOf(Zero)
	//显示反射类型对象的名称和种类
	fmt.Println(typeOfZero.Name(), typeOfZero.Kind())
	//Student struct
	//	Enum int
}
 */

/*
type Student struct {
	Name string
	Age  int
}
func main() {
	var stu Student
	typeOfStu := reflect.TypeOf(stu)
	fmt.Println(typeOfStu.Name(), typeOfStu.Kind())
//	Student struct
}
*/
