package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("a.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // 顺序有关
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	content := []byte("Go is an open source programing language that makes it easy to build simple,reliable,and efficient software")
	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
	}
}
