package main

import "fmt"

func main() {
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("quit")
				return
			default:
				fmt.Println("do something")
				// do something
			}
		}
	}()

	fmt.Println("main")
	quit <- true
}

// 主协程会早一步终止
