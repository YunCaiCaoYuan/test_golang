package main

import (
	"fmt"
	"strconv"
)

func main() {
	ret, _ := strconv.ParseInt("null", 10, 64)
	fmt.Println(ret)

	/*
		testMap := make(map[int32]bool)
		testMap[1] = true

		if testMap[2] {
			fmt.Println("found")
		}
		fmt.Println("not found")
		fmt.Println("testMap[2]", testMap[2])
	*/
}
