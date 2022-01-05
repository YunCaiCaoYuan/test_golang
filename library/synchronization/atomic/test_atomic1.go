package main

import (
	"fmt"
	"sync/atomic"
)

const x int64 = 1 + 1<<33

func main() {
	var i = x
	_ = i


	c := 4
	atomic.AddInt64(&i, ^int64(c-1))
	fmt.Println("^int64(c-1):", ^int64(c-1))


	fmt.Println("^int64(0):", ^int64(0))
	atomic.AddInt64(&i, ^int64(0)) // 取反
}
