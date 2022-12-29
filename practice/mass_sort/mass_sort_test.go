package main

import (
	"fmt"
	"testing"
)

func deleteElem(list []int, idx int) []int {
	return append(list[0:idx], list[idx+1:]...)
}

func printArr(list []int) {
	for _, item := range list {
		fmt.Printf("%d ", item)
	}
	fmt.Println("")
}

func Test_deleteElem(t *testing.T) {
	for i := 0; i < 5; i++ {
		list := []int{1, 2, 3, 4, 5}
		printArr(deleteElem(list, i))
	}
}
