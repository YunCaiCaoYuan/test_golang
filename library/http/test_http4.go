package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	fmt.Println("Method: ", r.Method)
	fmt.Println("URL: ", r.URL)
	fmt.Println("header: ", r.Header)
	fmt.Println("body: ", r.Body)
	fmt.Println("RemoteAddr: ", r.RemoteAddr)
	w.Write([]byte("请求成功!!!"))
}

func main() {

	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}
}
