package main

import "time"
import "fmt"

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

	/*
	Time, _ := time.Parse("2006-01-02 15:04:05", "2021-03-23 00:08:00")
	fmt.Println("Time=", Time)
	fmt.Println("Now=", time.Now())
	 */

	loc, _ := time.LoadLocation("Europe/Berlin")
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t)

	/*
	now := time.Now().Unix()
	fmt.Println("now=", now)
	 */

	/*
	now := time.Now().UnixNano()
	fmt.Println("now=", now) //  1619320725105078000
	 */

	/*
	nowStr := fmt.Sprintf("%d", time.Now().UnixNano())
	fmt.Println("nowStr=", nowStr) //  1619320725105078000
	 */

	/*
	timeUsed := time.Duration(148104536007)
	fmt.Println("timeUsed=", timeUsed)
	*/
}
