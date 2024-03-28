package main

import "fmt"

func main() {
	arr := make([]int, 0)
	for i := 0; i < 2000; i++ {
		fmt.Println("len为", len(arr), "cap为", cap(arr))
		arr = append(arr, i)
	}
}
