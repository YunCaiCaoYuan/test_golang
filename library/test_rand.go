package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(1234567890) // 初始化随机种子
	for i := 0; i<10; i++ {
		r := rand.Intn(10)
		fmt.Println(r)
	}
}
