package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

// net/http post demo

func main() {
	fmt.Println("start:", time.Now())
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=小王子&age=18"
	// json
	contentType := "application/json"
	data := `{"name":"小王子","age":18}`
	fmt.Println("http client...")
	var client = http.Client{
		Timeout: 1 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   1 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 8,
			IdleConnTimeout:     120 * time.Second,
			//DisableKeepAlives: true,
		},
	}
	//var client = http.Client{
		//Timeout: time.Second * 10, // net/http: request canceled (Client.Timeout exceeded while awaiting headers)
	//}
	resp, err := client.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Println("post failed, err=", err)
		fmt.Println("end:", time.Now())
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	fmt.Println("http client close...")
}
