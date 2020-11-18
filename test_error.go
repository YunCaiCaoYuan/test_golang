package main

import (
	"fmt"
)

func main() {
	fmt.Println("外层开始")
	defer func() {
		fmt.Println("外层准备recover")
		if err := recover(); err != nil {
			fmt.Printf("%#v-%#v\n", "外层", err) // err已经在上一级的函数中捕获了，这里没有异常，只是例行先执行defer，然后执行后面的代码
		} else {
			fmt.Println("外层没做啥事")
		}
		fmt.Println("外层完成recover")
	}()
	fmt.Println("外层即将异常")
	f()
	fmt.Println("外层异常后")
	defer func() {
		fmt.Println("外层异常后defer")
	}()
}

func f() {
	fmt.Println("内层开始")
	defer func() {
		fmt.Println("内层recover前的defer")
	}()

	defer func() {
		fmt.Println("内层准备recover")
		if err := recover(); err != nil {
			fmt.Printf("%#v-%#v\n", "内层", err) // 这里err就是panic传入的内容
		}

		fmt.Println("内层完成recover")
	}()

	defer func() {
		fmt.Println("内层异常前recover后的defer")
	}()

	panic("异常信息")

	defer func() {
		fmt.Println("内层异常后的defer")
	}()

	fmt.Println("内层异常后语句") //recover捕获的一级或者完全不捕获这里开始下面代码不会再执行
}
