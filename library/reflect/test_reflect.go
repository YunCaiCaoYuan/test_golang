package main

import "reflect"
import "fmt"

type MyMath struct {
	Pi float64
}

//普通函数
func (myMath MyMath) Sum(a, b int) int {
	return a + b
}

func (myMath MyMath) Dec(a, b int) int {
	return a - b
}

func main() {
	var myMath = MyMath{Pi:3.14159}
	//获取myMath的值对象
	rValue := reflect.ValueOf(myMath)
	//获取到该结构体有多少个方法
	//numOfMethod := rValue.NumMethod()
	//构造函数参数，传入两个整形值
	paramList := []reflect.Value{reflect.ValueOf(30), reflect.ValueOf(20)}

	//调用结构体的第一个方法Method(0)
	//注意:在反射值对象中方法索引的顺序并不是结构体方法定义的先后顺序
	//而是根据方法的ASCII码值来从小到大排序，所以Dec排在第一个，也就是Method(0)
	result := rValue.Method(0).Call(paramList)
	fmt.Println(result[0].Int())

	//	10
}
/*
//普通函数
func add(a, b int) int {
	return a + b
}
func main() {
	//将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)
	//构造函数参数，传入两个整形值
	paramList := []reflect.Value{reflect.ValueOf(2), reflect.ValueOf(3)}
	//反射调用函数
	retList := funcValue.Call(paramList)
	fmt.Println(retList[0].Int())
	// 5
}
 */

/*
func main() {
	type dog struct {
		LegCount int
	}
	//获取dog实例的反射值对象
	valueOfDog := reflect.ValueOf(&dog{})
	//// 取出dog实例地址的元素
	valueOfDog = valueOfDog.Elem()
	//获取legCount字段的值
	vLegCount := valueOfDog.FieldByName("LegCount")
	//尝试设置legCount的值
	vLegCount.SetInt(4)
	fmt.Println(vLegCount.Int())
	// 4
}
 */

/*
func main() {
	type dog struct {
		legCount int
	}
	//获取dog实例的反射值对象
	valueOfDog := reflect.ValueOf(&dog{})
	valueOfDog = valueOfDog.Elem()
	//获取legCount字段的值
	vLegCount := valueOfDog.FieldByName("legCount")
	//尝试设置legCount的值(这里会发生崩溃)
	vLegCount.SetInt(4)

//panic: reflect: reflect.flag.mustBeAssignable using value obtained using unexported field
}
 */

/*
func main() {
	//声明整形变量a并赋初值
	var a int = 1024
	//获取变量a的反射值对象
	rValue := reflect.ValueOf(&a)
	//取出a地址的元素(a的值)
	rValue = rValue.Elem()
	//尝试将a修改为1
	rValue.SetInt(1)
	//打印a的值
	fmt.Println(rValue.Int())

	//1
}
 */

/*
func main() {
	//声明整形变量a并赋初值
	var a int = 1024
	//获取变量a的反射值对象
	rValue := reflect.ValueOf(a)
	//尝试将a修改为1(此处会崩溃)
	rValue.SetInt(1)

	// panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
}
 */

/*
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
 */

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
