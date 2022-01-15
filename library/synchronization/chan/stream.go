package main

import (
	"fmt"
)

func asStream(done <-chan struct{}, values ...interface{}) <-chan interface{} {
	s := make(chan interface{}) //创建一个unbuffered的channel
	go func() { // 启动一个goroutine，往s中塞数据
		defer close(s) // 退出时关闭chan
		for _, v := range values { // 遍历数组
			select {
			case <-done:
				return
			case s <- v: // 将数组元素塞入到chan中
			}
		}
	}()
	return s
}


func takeN(done <-chan struct{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{}) // 创建输出流
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ { // 只读取前num个元素
			select {
			case <-done:
				return
			case takeStream <- <-valueStream: //从输入流中读取元素
			}
		}
	}()
	return takeStream
}

func main() {
	fmt.Println("asStream:")
	done := make(chan struct{})
	ch := asStream(done, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("takeN:")
	done = make(chan struct{})
	for v := range takeN(done, asStream(done, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10), 6) {
		fmt.Println(v)
	}

}

