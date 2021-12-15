package main

import (
	"fmt"
	"runtime"
	"sync"
)

var count int

func Count(lock *sync.Mutex) {
	lock.Lock() // 加写锁
	count++
	fmt.Println(count)
	lock.Unlock() // 解写锁，任何一个Lock()或RLock()均需要保证对应有Unlock()或RUnlock()
}

func main() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock) // 传递指针是为了防止函数内的锁和调用锁不一致
	}
	for {
		lock.Lock()
		c := count
		lock.Unlock()
		runtime.Gosched() // 交出时间片给协程
		if c > 10 {
			break
		}
	}
}
