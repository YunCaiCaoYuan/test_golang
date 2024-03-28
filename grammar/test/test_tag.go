package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	wg.Add(1)
	var test int
	// test := 1

	go func() {
		defer wg.Done()
	tag:
		fmt.Println("tag")

		if test > 0 {
			goto tag
		}

	}()

	wg.Wait()
}

// 会执行一次
// tag
