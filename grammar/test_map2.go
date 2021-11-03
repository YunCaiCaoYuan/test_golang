package main

import (
	"fmt"
	"sync"
)

// fatal error: concurrent map writes

func main() {
	fmt.Println("start...")
	testMap := make(map[int]int)
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i ++ {
		testMap[i] = i
	}


	for i := 0; i < 10000; i ++ {
		wg.Add(1)
		go func(key int) {
			delete(testMap, key)
		}(i)
		wg.Done()
	}

	wg.Wait()
	fmt.Println("finish...")
}
