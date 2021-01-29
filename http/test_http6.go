package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	resp, err := http.Get("http://127.0.0.1:9000")
	if err != nil {
		fmt.Printf("http.Get()函数执行错误,错误为:%v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("ioutil.ReadAll()函数执行出错,错误为:%v\n", err)
		return
	}

	fmt.Println(string(body))
}
