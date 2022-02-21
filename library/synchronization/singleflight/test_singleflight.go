package main

import (
	"log"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

var (
	sf           = singleflight.Group{}
	requestCount = int64(0)
	resp         = make(chan int64, 0)
	wg           sync.WaitGroup
)

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		//sf := singleflight.Group{}
		go func(v int) {
			log.Println("i123:", v)
			do, err, _ := sf.Do("number", Request)
			if err != nil {
				log.Println(err)
			}
			log.Println("resp:", do, "i:", v)
			defer wg.Done()
		}(i)
	}
	time.Sleep(time.Second)
	resp <- atomic.LoadInt64(&requestCount)
	wg.Wait()

}

func Request() (interface{}, error) {
	print("Request\n")
	atomic.AddInt64(&requestCount, 1)
	return <-resp, nil
}
