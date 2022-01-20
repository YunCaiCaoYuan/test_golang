package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Semaphore 数据结构，并且还实现了Locker接口
type semaphore struct {
	sync.Locker
	ch chan struct{}
}

// 创建一个新的信号量
func NewSemaphore(capacity int) sync.Locker {
	if capacity <= 0 {
		capacity = 1 // 容量为1就变成了一个互斥锁
	}
	return &semaphore{ch: make(chan struct{}, capacity)}
}

// 请求一个资源
func (s *semaphore) Lock() {
	s.ch <- struct{}{}
}

// 释放资源
func (s *semaphore) Unlock() {
	<-s.ch
}

var (
	maxWorkers = runtime.GOMAXPROCS(0)     // worker数量
	sema       = NewSemaphore(maxWorkers)  //信号量
	task       = make([]int, maxWorkers*4) // 任务数，是worker的四倍
)

func main() {
	print("maxWorkers123:", maxWorkers, "\n")
	print("task:", len(task), "\n")

	var wg sync.WaitGroup
	wg.Add(maxWorkers * 4)

	for i := range task {
		// 如果没有worker可用，会阻塞在这里，直到某个worker被释放
		sema.Lock()

		// 启动worker goroutine
		go func(i int) {
			defer sema.Unlock()
			defer wg.Done()
			time.Sleep(100 * time.Millisecond) // 模拟一个耗时操作
			task[i] = i + 1
		}(i)
	}

	wg.Wait()
	fmt.Println(task)
}
