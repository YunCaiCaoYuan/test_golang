package main

import "fmt"

type object struct {
	val int
}

func main() {

	testMap := make(map[int]*object)
	fmt.Println("len:", len(testMap))

	obj1 := &object{val : 123}
	testMap[123] = obj1
	fmt.Println("len:", len(testMap))
	//fmt.Println("before del, obj1=", obj1, "testMap=", testMap)

	delete(testMap, 123)
	//fmt.Println("after  del, obj1=", obj1, "testMap=", testMap)
}
