package main

import (
	"fmt"
	"sync"
)

// 等待所有 goroutine 执行完毕
// 进入死锁
func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})

	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt(i, done, wg)
	}

	close(done)
	wg.Wait()
	fmt.Println("all done!")
}

func doIt(workerID int, done <-chan struct{}, wg sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerID)
	defer wg.Done()
	<-done
	fmt.Printf("[%v] is done\n", workerID)
}
