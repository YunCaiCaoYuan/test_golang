package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}

func main() {
	var handler MyHandler
	var server = http.Server{
		Addr:           ":8080",
		Handler:        &handler,
		ReadTimeout:    2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	var err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("http server failed, err: %v\n", err)
		return
	}

	//resp, err := http.Get("http://example.com/")

	//resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)

	//resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})

	/*
		resp, err := http.Get("http://example.com/")
		if err != nil {
			// handle error
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
	*/
}
