package main

import "fmt"

func main() {
	var i = 1
	defer fmt.Println("result: ", func() int { return i * 2 }())
	i++
}
