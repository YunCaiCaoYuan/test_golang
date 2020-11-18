package main

import "fmt"

type test struct {
	Count int
	Price int
}

func main() {
	var testList []*test
	if testList == nil {
		fmt.Println("1111")
	}
}
