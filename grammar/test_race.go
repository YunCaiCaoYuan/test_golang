package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

/*
func main(){
	var x int
	go func(){
		for{
			x=1
		}
	}()

	go func(){
		for{
			x=2
		}
	}()
	time.Sleep(100*time.Second)
}
 */
// 一般并发的bug 有两种，死锁（block）和 竞争（race）
//
//死锁发生时，go run 会直接报错
//race 发生时，要加race 才会在运行时报warning
