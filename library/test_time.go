package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	nowStr := time.Now()
	fmt.Println("nowStr=", nowStr)
	nowStr2 := time.Now().Format("2006-01-02")
	fmt.Println("nowStr2=", nowStr2)
	nowStr3 := time.Now().Unix()
	fmt.Println("nowStr3=", nowStr3)
	 */

	/*
	timeAdd := time.Minute+5*time.Second
	fmt.Println("timeAdd=", timeAdd)
	println("uint64 timeAdd=", timeAdd)
	 */

	now := time.Now().Unix()
	fmt.Println("now=", now)


	/*
	timeUsed := time.Duration(148104536007)
	fmt.Println("timeUsed=", timeUsed)
	*/
}
