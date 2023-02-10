package main

import (
	"arena"
	"fmt"
)

type T struct {
	val int
}

func main() {
	are := arena.NewArena()
	t := arena.New[T](are)
	t.val = 100
	fmt.Println(t)
	are.Free()

	// free 后会panic
	t = arena.New[T](are)
	t.val = 100
	fmt.Println(t)
}
