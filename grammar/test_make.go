package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type test struct {
	Count int
	Price int
}

func (u *test) TableName() string {
	return "channel_skip_id"
}

func returnSlice() []int32 {
	list := make([]int32, 0)
	var v1,v2 int32 = 1,2
	list = append(list, v1, v2)
	fmt.Printf("list1=%p\n", &list)
	fmt.Println("list1=", list)
	fmt.Printf("list1 =%p\n", &list[0])
	return list
}

func returnStructSlice() []*test {
	list := make([]*test, 0, 2)
	list = append(list, &test{Count: 1},&test{Count: 2})
	sh:= (*reflect.SliceHeader)(unsafe.Pointer(&list))
	fmt.Printf("list1=%p\n", list)
	fmt.Println("list1=", list)
	fmt.Printf("list1 =%p\n", &list[0])
	fmt.Println("list1 0=", sh)
	return list
}

func main() {
	list := returnStructSlice() // 上层结构浅拷贝，底层地址不变
	fmt.Printf("list2=%p\n", &list)
	fmt.Println("list2=", list)
	fmt.Printf("list2=%p\n", &list[0])
	//list1=0xc00000c0a0
	//list1= [0xc00001a090 0xc00001a0a0]
	//list2=0xc00000c080
	//list2= [0xc00001a090 0xc00001a0a0]


	/*
	const test = 4
	println(reflect.TypeOf(test).Name()) // int
	 */

	// 测试不用make给切片分配内存，testList=nil
	//var testList []*test
	//fmt.Println("testList-len=", len(testList)) // len=0
	//if testList == nil {
	//	fmt.Println("1111")
	//}

	// 测试var定义结构体，是否可以调用方法，可以
	//var testVar test
	//testVar.Count = 10
	//fmt.Println("TableName=", testVar.TableName())
	//fmt.Println("testVar.Count=", testVar.Count)

	// 测试切片长度，len=0
	//testList := make([]*test, 0)
	//fmt.Println("testList.len=", len(testList))
	//fmt.Println("testList.len=", len(nil))

	// test switch
	//num := 1
	//switch num {
	//case 1	:
	//	//fmt.Println("1")
	//case 2	:
	//	fmt.Println("2")
	//default:
	//	fmt.Println("default")
	//}

	// test 格式化
	//sql := fmt.Sprintf("head '%s' num '%d' ", "a string ", 10000)
	//fmt.Println("sql : ", sql)
	//
	//sql2 := fmt.Sprintf("head %s num %d ", "a string ", 10000)
	//fmt.Println("sql2 : ", sql2)

	//Emoji
	//r := regexp.MustCompile("^[ a-zA-Z0-9\u4e00-\u9fa5\u1F300-\u1F5FF\u1F600-\u1F64F\u1F680-\u1F6FF\u2600-\u2B55]{1,4}$")
	//name := "ἰ0"
	//if r.MatchString(name) {
	//	fmt.Println("ok!!!")
	//}
}
