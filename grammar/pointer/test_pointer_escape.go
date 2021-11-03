package main

//import "fmt"

type object struct {
	val int
}

func main() {
	newObject(123)
	//fmt.Printf("obj1=%p\n", obj1)
	//fmt.Println("obj1=", obj1)
}

func newObject(val int) *object {
	funcObj1 := &object{
		val: val,
	}
	//fmt.Printf("obj0=%p\n", funcObj1)
	return funcObj1
}
// 逃逸分析：
/*
./test_pointer_escape.go:15:6: can inline newObject
./test_pointer_escape.go:9:6: can inline main
./test_pointer_escape.go:10:11: inlining call to newObject
/var/folders/6v/tmrxzp7s6jdfp60spw7ljcyc0000gn/T/go-build030433014/b001/_gomod_.go:6:6: can inline init.0
./test_pointer_escape.go:10:11: main &object literal does not escape
./test_pointer_escape.go:17:3: &object literal escapes to heap
*/
