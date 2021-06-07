package main

import (
	"fmt"
	"net/http"
	"sync"
)

var table sync.Map

type Getter func() *http.Request

func GetRequest(url string) *http.Request {
	getter := getReqFromMap(url)
	// 從getter裡取出真正的Request
	return getter()
}

func getReqFromMap(url string) Getter {
	if f, ok := table.Load(url); ok {
		return f.(Getter)
	}

	// 每個Load找不到的goroutine可能同時執行以下這段程式
	var req *http.Request
	var wg sync.WaitGroup

	wg.Add(1)
	waitGetter := func() *http.Request {
		wg.Wait()
		return req
	}

	f, loaded := table.LoadOrStore(url, Getter(waitGetter))
	if loaded {
		return f.(Getter)
	}
	fmt.Println("loaded=", loaded)

	// Store成功，初始化這個Request
	req, _ = http.NewRequest(http.MethodGet, url, nil)

	// 通知其他goroutine這個req已經建立完成
	wg.Done()

	// 把Getter換成沒有Wait()的版本，利於GC和加快速度
	wrapGetter := func() *http.Request {
		return req
	}
	table.Store(url, Getter(wrapGetter))
	return Getter(wrapGetter)
}

func main() {
	var wg sync.WaitGroup
	var req *http.Request

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			req = GetRequest("http://example.com/user/")
		}(i)
	}
	wg.Wait()
}
