package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
	"time"
)

func main() {
	var (
		sf = singleflight.Group{}
		wg sync.WaitGroup
	)

	//intMap := make(map[string]int)
	var intMap sync.Map

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(v int) {
			defer wg.Done()

			key := "abc"
			//if val, ok := intMap[key]; !ok {
			if val, ok := intMap.Load(key); !ok {

				do, err, shared := sf.Do(key, func() (interface{}, error) {
					fmt.Println("mock exec, no val:", "vth:", v)

					// 提高下游请求并发
					go func() {
						time.Sleep(10 * time.Millisecond)
						fmt.Printf("Deleting key: %v, v: %v\n", key, v)
						sf.Forget(key)
					}()

					// fixme db操作或耗时操作
					time.Sleep(15 * time.Millisecond)
					return 123, nil
				})

				if err != nil {
					fmt.Println("err:", err)
					return
				}

				if newVal, ok := do.(int); ok {
					//intMap[key] = newVal
					intMap.Store(key, newVal)
					fmt.Println("singleflight val:", "vth:", v, "newVal:", newVal, "shared:", shared)
				}

			} else {
				fmt.Println("read memory, val:", val, "vth:", v)
			}
		}(i)

		// 模拟上游请求: 分多次并发
		if i == 50 {
			time.Sleep(10 * time.Millisecond)
		}
	}

	wg.Wait()
}

// 用map，会出现并发写:
// fatal error: concurrent map writes
// 不是必现，大概率出现

// 用sync.map和singleflight
// sync.map保证并发安全，singleflight合并并发db操作，降低请求数量级。
