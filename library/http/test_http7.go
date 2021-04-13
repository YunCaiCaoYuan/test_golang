package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	//1.处理请求参数
	params := url.Values{}
	params.Set("name", "itbsl")
	params.Set("hobby", "fishing")

	//2.设置请求URL
	rawUrl := "http://127.0.0.1:9000"
	reqURL, err := url.ParseRequestURI(rawUrl)
	if err != nil {
		fmt.Printf("url.ParseRequestURI()函数执行错误,错误为:%v\n", err)
		return
	}

	//3.整合请求URL和参数
	//Encode方法将请求参数编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序。
	reqURL.RawQuery = params.Encode()

	//4.发送HTTP请求
	//说明: reqURL.String() String将URL重构为一个合法URL字符串。
	resp, err := http.Get(reqURL.String())
	if err != nil {
		fmt.Printf("http.Get()函数执行错误,错误为:%v\n", err)
		return
	}
	defer resp.Body.Close()

	//5.一次性读取响应的所有内容
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("ioutil.ReadAll()函数执行出错,错误为:%v\n", err)
		return
	}

	fmt.Println(string(body))
}
