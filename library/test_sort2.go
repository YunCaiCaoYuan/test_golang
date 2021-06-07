package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []int32{4, 40, 2, 20, 1, 10}
	sort.Slice(list, func(i, j int) bool {
		return list[i] > list[j]
	}) // i:1 j:0
	fmt.Println("list=", list)
}
