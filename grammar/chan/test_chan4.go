package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, (i+1)*100)
	}
}
func say2(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s, (i+1)*150)
	}
}
func main() {
	go say2("world")
	say("hello")
}
