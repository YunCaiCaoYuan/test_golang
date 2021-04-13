package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func checkError(err error) {
	if err != nil{
		log.Fatalln(err)
	}
}

func main() {
	tr := http.Transport{DisableKeepAlives: true}
	client := http.Client{Transport: &tr}

	resp, err := client.Get("https://golang.google.cn/")
	if resp != nil {
		defer resp.Body.Close()
	}
	checkError(err)

	fmt.Println(resp.StatusCode)    // 200

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(len(string(body)))
}
/*
// 主动关闭连接
func main() {
	//req, err := http.NewRequest("GET", "http://golang.org", nil)
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	checkError(err)

	req.Close = true
	//req.Header.Add("Connection", "close")    // 等效的关闭方式

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	checkError(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fmt.Println(string(body))
}
*/
