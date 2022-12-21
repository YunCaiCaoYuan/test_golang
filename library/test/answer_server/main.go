package main

import (
	"fmt"
	"net/http"
)

func handleCommitAnswer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path:", r.URL.Path, "method:", r.Method)
	if r.Method == "GET" {
		w.Write([]byte("hello1"))
		return
	}

	if r.Method == "POST" {
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
