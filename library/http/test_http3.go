package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

//实现Handler接口
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "goodbye")
}

func main() {
	var handler MyHandler
	http.Handle("/sayGoodbye", &handler)
	var err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server failed, err: %v\n", err)
		return
	}
}
