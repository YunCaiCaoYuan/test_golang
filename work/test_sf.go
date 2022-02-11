package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"sync"
)

var (
	sf = singleflight.Group{}
	wg sync.WaitGroup
)

func main() {
	//intMap := make(map[string]int)
	var intMap sync.Map

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func(v int) {
			defer wg.Done()

			key := "abc"
			//if val, ok := intMap[key]; !ok {
			if val, ok := intMap.Load(key); !ok {

				do, err, _ := sf.Do(key, func() (interface{}, error) {
					// fixme db操作或耗时操作
					fmt.Println("no val:", "vth:", v)
					return 123, nil
				})
				if err != nil {
					fmt.Println("err:", err)
					return
				}

				if newVal, ok := do.(int); ok {
					//intMap[key] = newVal
					intMap.Store(key, newVal)
					fmt.Println("create val:", "vth:", v)
				}

			} else {
				fmt.Println("val:", val, "vth:", v)
			}
		}(i)
	}

	wg.Wait()
}

// 用map，会出现并发写:
// fatal error: concurrent map writes
// 不是必现，大概率出现

// 用sync.map和singleflight
// sync.map保证并发安全，singleflight合并并发db操作，保证load一次db。
