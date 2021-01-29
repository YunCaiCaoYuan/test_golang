package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm) // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func main() {
	http.HandleFunc("/", postHandler)
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}

	//client := &http.Client{
	//	CheckRedirect: redirectPolicyFunc,
	//}
	//resp, err := client.Get("http://example.com")

	//req, err := http.NewRequest("GET", "http://example.com", nil)

	//req.Header.Add("If-None-Match", `W/"wyzzy"`)
	//resp, err := client.Do(req)

	/*
		tr := &http.Transport{
			TLSClientConfig:    &tls.Config{RootCAs: pool},
			DisableCompression: true,
		}
		client := &http.Client{Transport: tr}
		resp, err := client.Get("https://example.com")
	*/
}
