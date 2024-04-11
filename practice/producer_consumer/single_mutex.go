package main

import (
	"fmt"
	"sync"
)

const BufferSize = 5

var buffer []int
var mutex sync.Mutex
var cond = sync.NewCond(&mutex)

func producer(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		for len(buffer) == BufferSize {
			cond.Wait()
		}
		buffer = append(buffer, i)
		fmt.Printf("producer: %d\n", i)
		cond.Signal()
		mutex.Unlock()
	}
	wg.Done()
}

func consumer(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		mutex.Lock()
		for len(buffer) == 0 {
			cond.Wait()
		}
		item := buffer[0]
		buffer = buffer[1:]
		fmt.Printf("consumer:%d\n", item)
		cond.Signal()
		mutex.Unlock()
	}
	wg.Done()
}

func main() {
	buffer = make([]int, 0, BufferSize)

	var wg sync.WaitGroup
	wg.Add(2)

	go producer(&wg)
	go consumer(&wg)

	wg.Wait()
	fmt.Println("All done")
}
