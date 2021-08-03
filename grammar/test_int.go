package main

import "fmt"

func main() {
	fmt.Println("0 / 10 =", 0 / 10 * 10)
	fmt.Println("9 / 10 =", 9 / 10 * 10)
	fmt.Println("12 / 10 =", 12 / 10 * 10)
	fmt.Println("20 / 10 =", 20 / 10 * 10)

	var intval int64
	fmt.Println("intval=", intval)

	/*
	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	fmt.Println(unsafe.Sizeof(i1)) // 8
	fmt.Println(unsafe.Sizeof(i2)) // 1
	fmt.Println(unsafe.Sizeof(i3)) // 2
	fmt.Println(unsafe.Sizeof(i4)) // 4
	fmt.Println(unsafe.Sizeof(i5)) // 8
	 */
}
