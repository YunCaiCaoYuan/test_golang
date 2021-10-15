package main

import (
	"sync"
	"time"
)

func main() {
	mu := sync.RWMutex{}
	mu.RLock()
	defer mu.RUnlock() // do not unlock in defer
	go func(mu *sync.RWMutex) {
		mu.Lock()
		defer mu.Unlock()
	}(&mu)
	time.Sleep(time.Second) // a rough way to ensure new goroutine running
	mu.RLock() // deadlock
	defer mu.RUnlock()
}
