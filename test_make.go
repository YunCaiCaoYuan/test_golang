package main

import "fmt"

type test struct {
	Count int
	Price int
}

func (u *test) TableName() string {
	return "channel_skip_id"
}

func main() {
	// 测试不用make给切片分配内存
	//var testList []*test
	//if testList == nil {
	//	fmt.Println("1111")
	//}

	// 测试var定义结构体，是否可以调用方法
	var testVar test
	testVar.Count = 10
	fmt.Println("TableName=", testVar.TableName())
	fmt.Println("testVar.Count=", testVar.Count)
}
