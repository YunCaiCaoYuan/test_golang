package main

import (
	"runtime"
	"time"
)

var _ = runtime.GOMAXPROCS(3)

var a, b int

func u1() {
	a = 1
	b = 2
}

func u2() {
	a = 3
	b = 4
}

func p() {
	println(a)
	println(b)
}

func main() {
	go u1()    // 多个 goroutine 的执行顺序不定
	go u2()
	go p()
	time.Sleep(1 * time.Second)
}
