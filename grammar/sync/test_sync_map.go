package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var table sync.Map

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if _, ok := table.Load("KEY"); !ok {
				table.Store("KEY", n)
			}
		}(i)
	}
	wg.Wait()
	val, ok := table.Load("KEY")
	fmt.Println(val, ok)
}
// 這段程式的執行結果可能print出0~99間任意的val值，而非固定是0。
