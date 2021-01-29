package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := r.URL.Query()
	fmt.Fprintln(w, "name:", params.Get("name"), "hobby:", params.Get("hobby"))
}

func main() {

	http.HandleFunc("/", myHandler)
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		fmt.Printf("http.ListenAndServe()函数执行错误,错误为:%v\n", err)
		return
	}
}
