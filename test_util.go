package main

import "fmt"

func main() {
	testMap := make(map[int32]bool)
	testMap[1] = true

	if testMap[2] {
		fmt.Println("found")
	}
	fmt.Println("not found")
	fmt.Println("testMap[2]", testMap[2])
}
