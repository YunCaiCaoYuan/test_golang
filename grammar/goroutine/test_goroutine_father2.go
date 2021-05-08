package main

import "time"
import "fmt"
import "context"

func main() {
	fmt.Println("main 函数 开始...")
	go func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		fmt.Println("父 协程 开始...")
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("子 协程 接受停止信号...")
					return
				default:
					fmt.Println("子 协程 执行中...")
					timer := time.NewTimer(time.Second * 2)
					<-timer.C
				}
			}
		}(ctx)
		time.Sleep(time.Second*5)
		fmt.Println("父 协程 退出...")
	}()
	time.Sleep(time.Second*10)
	fmt.Println("main 函数 退出")
}
