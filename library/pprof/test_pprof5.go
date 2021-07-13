package main

import (
	"log"
	"net/http"
	//引入 提供采集数据http接口的包
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

const (
	Ki = 1024
	Mi = Ki * Ki
	Gi = Ki * Mi
)

func main() {
	runtime.GOMAXPROCS(1)              // 限制 CPU 使用数，避免过载
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪

	go func() {
		// 开启http服务
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	for {
		eatCpu()
		stealMemory()
		time.Sleep(time.Second)
	}
}

func eatCpu() {
	loop := 10000000000
	for i := 0; i < loop; i++ {
		// do nothing
	}
}

func stealMemory() {
	max := Gi
	var buffer [][Mi]byte
	// 每次向buffer中 追加 1MB，直到1GB为止
	for len(buffer)*Mi < max {
		buffer = append(buffer, [Mi]byte{})
	}
}
