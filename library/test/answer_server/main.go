package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type param struct {
	TitleId   int `json:"title_id"`
	AnswerOpt int `json:"answer_opt"`
}

func handleCommitAnswer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path:", r.URL.Path, "method:", r.Method, "form:", r.Form, "PostForm:", r.PostForm, "header:", r.Header)
	if r.Method == "GET" {
		w.Write([]byte("hello1"))
		return
	}

	if r.Method == "POST" {
		b, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		fmt.Println("len(b):", len(b), "cap(b):", cap(b))
		fmt.Println("body:", string(b))

		arg := new(param)
		err := json.Unmarshal(b, arg)
		if err != nil {
			fmt.Println("err:", err.Error())
			return
		}
		fmt.Println("arg:", arg)

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
