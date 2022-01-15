package main

import (
	"fmt"
	"time"
)

// 大佬实现的
type Token struct{}

func newWorker(id int, ch chan Token, nextCh chan Token) {
	for {
		token := <-ch       // 取得令牌
		fmt.Println(id + 1) // id从1开始
		time.Sleep(time.Second)
		nextCh <- token
	}
}
func main() {
	chs := []chan Token{make(chan Token), make(chan Token), make(chan Token), make(chan Token)}

	// 创建4个worker
	for i := 0; i < 4; i++ {
		go newWorker(i, chs[i], chs[(i+1)%4])
	}

	//首先把令牌交给第一个worker
	chs[0] <- struct{}{}

	select {}
}

// 我实现的 击鼓传花
/*
var turn1 = make(chan int, 1)
var turn2 = make(chan int, 1)
var turn3 = make(chan int, 1)
var turn4 = make(chan int, 1)

func g1() {
	for {
		<- turn1
		fmt.Println("1")
		time.Sleep(time.Second)
		turn2 <- 1
	}
}
func g2() {
	for {
		<- turn2
		fmt.Println("2")
		time.Sleep(time.Second)
		turn3 <- 1
	}
}
func g3() {
	for {
		<- turn3
		fmt.Println("3")
		time.Sleep(time.Second)
		turn4 <- 1
	}
}
func g4() {
	for {
		<- turn4
		fmt.Println("4")
		time.Sleep(time.Second)
		turn1 <- 1
	}
}

func main() {
	turn1 <- 1
	go g1()
	go g2()
	go g3()
	go g4()
	select {}
}
*/
