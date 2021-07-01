package main

import "fmt"
import "reflect"

func main() {
	//*int的空指针
	var a *int
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())
	//nil值
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid())
	//*int类型的空指针
	fmt.Println("(*int)(nil):", reflect.ValueOf((*int)(nil)).Elem().IsValid())
	//实例化一个结构体
	s := struct {}{}
	//尝试从结构体中查找一个不存在的字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(s).FieldByName("").IsValid())
	//尝试从结构体中查找一个不存在的方法
	fmt.Println("不存在的方法:", reflect.ValueOf(s).MethodByName("").IsValid())
	//实例化一个map
	m := map[int]int{}
	//尝试从map中查找一个不存在的键
	fmt.Println("不存在的键:", reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())

	//var a *int: true
	//nil: false
	//(*int)(nil): false
	//不存在的结构体成员: false
	//不存在的方法: false
	//不存在的键: false

}

/*
//定义结构体
type Student struct {
	Name string
	Age  int
	//嵌入字段
	float32
	bool
	next *Student
}
func main() {
	//值包装结构体
	rValue := reflect.ValueOf(Student{
		next: &Student{},
	})
	//获取字段数量
	fmt.Println("NumField:", rValue.NumField())

	//获取索引为2的字段(float32字段)
	//注:经过测试发现Field(i)的参数索引是从0开始的，
	//并且是按照定义的结构体的顺序来的，而不是按照字段名字的ASCii码值来的
	floatField := rValue.Field(2)
	//输出字段类型
	fmt.Println("Field:", floatField.Type())
	//根据名字查找字段
	fmt.Println("FieldByName(\"Age\").Type:", rValue.FieldByName("Age").Type())
	//根据索引查找值中next字段的int字段的值
	fmt.Println("FieldByIndex([]int{4, 0}).Type()", rValue.FieldByIndex([]int{4,0}).Type())

	//NumField: 5
	//Field: float32
	//FieldByName("Age").Type: int
	//FieldByIndex([]int{4, 0}).Type() string
}
 */

/*
type Bird struct {
	Name string
	LifeExpectance int
}
func (b * Bird) Fly() {
	fmt.Println("I am flying...")
}
func main() {
	sparrow := &Bird{"Sparrow", 3}
	S := reflect.ValueOf(sparrow).Elem()
	typeOfT := S.Type()
	for i := 0; i < S.NumField(); i++ {
		f := S.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
//	0: Name string = Sparrow
//  1: LifeExpectance int = 3
}
 */

