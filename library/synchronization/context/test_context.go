package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer func() {
			fmt.Println("goroutine exit")
		}()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("ctx.Done()")
				return
			default:
				fmt.Println("Sleep...")
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Second)
	cancel()
	time.Sleep(2 * time.Second)
}