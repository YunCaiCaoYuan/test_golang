package main

import "fmt"
import "reflect"

func main() {
	//声明整型变量a并赋初值
	var a int = 1024
	//获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)
	//获取interface{}类型的值，通过类型断言转换
	var getA int = valueOfA.Interface().(int)
	//获取64位的值，强制类型转换为int类型
	var getB int = int(valueOfA.Int())
	fmt.Println(getA, getB)
}

/*
func main() {
	// 声明一个空结构体
	type cat struct {
		Name string
		Type int `json:"type" id:"100"` // 带有结构体tag的字段
	}
	// 创建cat的实例
	ins := cat{Name: "mimi", Type: 1}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}
 */

/*
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
 */

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
