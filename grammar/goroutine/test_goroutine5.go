package main

import (
	"fmt"
	"time"
)

/*
func main() {
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- idx
		}(i)
	}

	fmt.Println(<-ch)        // 输出第一个发送的值
	close(ch)            // 不能关闭，还有其他的 sender
	time.Sleep(2 * time.Second)    // 模拟做其他的操作
}
 */
func main() {
	ch := make(chan int)
	done := make(chan struct{})

	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, "Send result")
			case <-done:
				fmt.Println(idx, "Exiting")
			}
		}(i)
	}

	fmt.Println("Result: ", <-ch)
	close(done)
	time.Sleep(3 * time.Second)
}
