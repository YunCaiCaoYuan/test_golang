package main

import "time"
import "fmt"

func main() {
	fmt.Println("main 函数 开始...")
	go func() {
		fmt.Println("父 协程 开始...")
		go func() {
			for {
				fmt.Println("子 协程 执行中...")
				timer := time.NewTimer(time.Second * 2)
				<-timer.C
			}
		}()
		time.Sleep(time.Second*5)
		fmt.Println("父 协程 退出...")
	}()
	time.Sleep(time.Second*10)
	fmt.Println("main 函数 退出")
}
