package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("生产者生产数据: ", i)
		time.Sleep(time.Second)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for {
		data, ok := <-ch
		if !ok {
			fmt.Println("消费者消费完数据，退出")
			break
		}

		fmt.Println("消费者消费数据: ", data)
		time.Sleep(2 * time.Second)
	}
}

func main() {
	ch := make(chan int, 3)
	go producer(ch)
	go consumer(ch)

	time.Sleep(20 * time.Second)
}

/*
生产者生产数据:  0
消费者消费数据:  0
生产者生产数据:  1
消费者消费数据:  1
生产者生产数据:  2
生产者生产数据:  3
消费者消费数据:  2
生产者生产数据:  4
消费者消费数据:  3
消费者消费数据:  4
消费者消费完数据，退出
*/
