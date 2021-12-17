package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//r1 := rand.Perm(7)
	//fmt.Println(r1)

	a := []int{0,1,2,3,4,5,6}
	for i:=0; i<10; i++ {
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		fmt.Println(a)
	}


	/*
	rand.Seed(1234567890) // 初始化随机种子
	for i := 0; i<10; i++ {
		r := rand.Intn(10)
		fmt.Println(r)
	}
	 */
}
