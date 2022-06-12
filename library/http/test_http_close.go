package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path", r.URL.Path)
	fmt.Fprintln(w, "hello world")
}

func main() {

	//1. 创建http请求处理器，并注册path与处理器
	handler := http.NewServeMux()
	handler.HandleFunc("/hello", sayHello)

	//2. 创建httpserver
	server := &http.Server{Addr: ":9090", Handler: handler}
	// 用于实现通知等待模型
	shutdownFlag := make(chan int)

	//3. 异步协程，模拟优雅停服
	go func() {
		//3.1 休眠1分钟，让服务器有1分钟时间可以处理请求
		time.Sleep(time.Minute)
		fmt.Println("call shutdown")
		//3.2如果1分钟还没关闭完毕，则报错
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		//3.3优雅关闭
		err1 := server.Shutdown(ctx)
		if err1 != nil {
			fmt.Println(err1)

		}

		//3.3通知main协程，关闭完毕，可以退出了
		close(shutdownFlag)
	}()

	//4. 启动http服务
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		panic(err)
	}

	//5. 等待优雅关闭结束
	<-shutdownFlag
	fmt.Println("shutdown ok ")
}
