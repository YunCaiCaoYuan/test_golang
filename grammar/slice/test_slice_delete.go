package main

import "fmt"

func main() {
	sliceLen := 1
	list := make([]int, 0, sliceLen)
	for i:=1; i<=sliceLen; i++{
		list = append(list, i)
	}
	fmt.Println("before del list:", list, "len:", len(list), "cap:", cap(list))

	delVal := 1
	delIdx := -1
	for idx, val := range list {
		if delVal == val {
			delIdx = idx
		}
	}
	fmt.Println("del val:", delVal, "del index:", delIdx)

	if delIdx > 0 {
		list = append(list[:delIdx], list[delIdx+1:]...)
	}
	fmt.Println("after del list ", list, "len:", len(list), "cap:", cap(list))

	list = list[:0]
	fmt.Println("after reset list ", list, "len:", len(list), "cap:", cap(list))
}
