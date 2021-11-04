package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // 初始化随机种子
	for i := 0; i<10; i++ {
		r := rand.Intn(10)
		fmt.Println(r)
	}
}
