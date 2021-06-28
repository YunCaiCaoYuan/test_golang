package main

import "fmt"

// break 配合 label 跳出指定代码块
/*
func main() {
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			//break    // 死循环，一直打印 breaking out...
			break loop
		}
	}
	fmt.Println("out...")
}
 */
/*
func main() {
	data := []string{"one", "two", "three"}

	for _, v := range data {
		go func() {
			fmt.Println(v)
		}()
	}

	time.Sleep(3 * time.Second)
	// 输出 three three three
}
 */

type testStruct struct {
	id int32
	name string
}
func returnList() []*testStruct {
	return nil
}
func main() {
	list := returnList()
	for _,item := range list {
		fmt.Println("item=", item)
	}
	fmt.Println(list)

	/*
	data := []string{"one", "two", "three"}

	for _, v := range data {
		go func(in string) {
			fmt.Println(in)
		}(v)
	}

	time.Sleep(3 * time.Second)
	 */
	// 输出 one two three
}
