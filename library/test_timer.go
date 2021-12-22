package main

import (
	"fmt"
	"time"
)

func main() {
	go timerTest()
	select {}
}

func timerTest() {
	ticker := time.NewTicker(time.Second*3)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick...", time.Now().Unix())
		}
	}
}
