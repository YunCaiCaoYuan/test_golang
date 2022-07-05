package main

import "fmt"

func test_e1() bool {
	fmt.Println("test_e1")
	return false
}

func test_e2() bool {
	fmt.Println("test_e2")
	return true
}

func main() {
	if test_e1() && test_e2() { // 没错是短路啊！！！
		fmt.Println("in main")
	}
}
