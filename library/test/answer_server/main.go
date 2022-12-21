package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleCommitAnswer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path:", r.URL.Path, "method:", r.Method, "form:", r.Form, "PostForm:", r.PostForm, "header:", r.Header)
	if r.Method == "GET" {
		w.Write([]byte("hello1"))
		return
	}

	if r.Method == "POST" {
		b, _ := ioutil.ReadAll(r.Body)
		fmt.Println("len(b):", len(b), "cap(b):", cap(b))
		fmt.Println("body:", string(b))

		//fmt.Println("r.Form.Get(passwd)", r.Form.Get("passwd"))

		//decoder := json.NewDecoder(r.Body)
		//var params map[string]string
		//decoder.Decode(&params)
		//fmt.Println("len(params):", len(params))
		//for k, v := range params {
		//	fmt.Println("k:", k, "v:", v)
		//}

		w.Write([]byte("hello2"))
		return
	}
}

func main() {
	fmt.Println("answer start...")

	handler := http.NewServeMux()
	handler.HandleFunc("/commit_answer", handleCommitAnswer)

	server := &http.Server{Addr: ":8080", Handler: handler}
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		panic(err)
	}
}
