package main

import (
	"fmt"
	"time"
)

const (
	NumProducers = 4
	NumConsumers = 2
	BufferSize   = 3
	NumItems     = 20
)

func getGoroutineID() int {
	return int(time.Now().UnixNano() / 1000000)
}

func producer(buffer chan<- int) {
	for i := 0; i < NumItems; i++ {
		buffer <- i
		fmt.Printf("Producer %d produced %d\n", getGoroutineID(), i)
	}
}

func consumer(buffer <-chan int) {
	for item := range buffer {
		fmt.Printf("Consumer %d consumed %d\n", getGoroutineID(), item)
	}
}

func main() {
	buffer := make(chan int, BufferSize)

	// 启动生产者
	for i := 0; i < NumProducers; i++ {
		go producer(buffer)
	}

	// 启动消费者
	for i := 0; i < NumConsumers; i++ {
		go consumer(buffer)
	}

	time.Sleep(10 * time.Second)
	close(buffer)
}
