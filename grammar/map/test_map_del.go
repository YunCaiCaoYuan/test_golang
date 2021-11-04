package main

type object struct {
	val int
}

func main() {

	testMap := make(map[int]*object)
	obj1 := &object{val : 123}
	testMap[123] = obj1
	//fmt.Println("before del, obj1=", obj1, "testMap=", testMap)

	delete(testMap, 123)
	//fmt.Println("after  del, obj1=", obj1, "testMap=", testMap)
}
