package main

func main() {
	// 1：无缓冲通道，发送方和接收方必须同时准备好，否则会阻塞
	// var ch chan int
	// ch <- 1

	// 2：有缓冲通道，发送方和接收方不需要同时准备好
	var ch = make(chan int, 1)
	ch <- 1
}

// fatal error: all goroutines are asleep - deadlock!
// 阻塞死锁
