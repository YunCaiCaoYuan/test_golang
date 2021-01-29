package main

import (
	"fmt"
	"net/http"
)

//say hello to the world
func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "hello world")
	//_, _ = w.Write([]byte("hello world"))
}

func main() {

	//1.注册一个处理器函数
	http.HandleFunc("/", sayHello)

	//2.设置监听的TCP地址并启动服务
	//参数1:TCP地址(IP+Port)
	//参数2:handler handler参数一般会设为nil，此时会使用DefaultServeMux。
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}
}
