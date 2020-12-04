package main

import (
	"fmt"
	"time"
)

func main() {
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("nowStr=", nowStr)

}
