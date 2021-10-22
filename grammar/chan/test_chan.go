package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}
}

type EventData struct {
	Event int
	Data int
}

func main() {
	myChan := make(chan EventData)

	/*
	go say("world")
	say("hello")
	 */
}
