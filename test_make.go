package main

import "fmt"

type test struct {
	Count int
	Price int
}

// 测试不用make给切片分配内存

func main() {
	var testList []*test
	if testList == nil {
		fmt.Println("1111")
	}
}
