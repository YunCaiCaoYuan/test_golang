package main

import "time"
import "fmt"

func GetChanStarUpgradeCycle(t time.Time) (begin time.Time, statisticEnd time.Time, cycleEnd time.Time) {
	year, month, _ := t.Date()
	var (
		splitDate = time.Date(year, month, 16, 0, 0, 0, 0, t.Location())
	)
	if t.Sub(splitDate) > 0 {
		if month == time.February {
			statisticEnd = time.Date(year, month, 27, 0, 0, 0, 0, t.Location())
		} else {
			statisticEnd = time.Date(year, month, 29, 0, 0, 0, 0, t.Location())
		}
		begin = time.Date(year, month, 16, 0, 0, 0, 0, t.Location())
		cycleEnd = time.Date(year, month+1, 1, 0, 0, 0, 0, t.Location()).Add(-1)
	} else {
		begin = time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
		statisticEnd = time.Date(year, month, 14, 0, 0, 0, 0, t.Location())
		cycleEnd = time.Date(year, month, 16, 0, 0, 0, 0, t.Location()).Add(-1)
	}

	return
}

func main() {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(now)

	/*
		birth := "2010-01-01"
		tim, err := time.Parse("2006-01-02", birth)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(tim)
	*/

	/*
		dur := time.Duration(-1) // 0
		ret := int64(dur / time.Second)
		fmt.Println("ret=", ret) // 0
	*/

	/*
		now := time.Now()
		format3 := now.Format("01月02日")
		fmt.Println("format3=", format3)
	*/

	//now := time.Now().Add(time.Duration(60) * time.Second)
	//fmt.Println("now=", now)
	//begin,_,_ := GetChanStarUpgradeCycle(now)
	//fmt.Println("begin=", begin)
	//begin_format := begin.Format("2006/01/02")
	//fmt.Println("begin_format=", begin_format)
	//format := now.Format("2006")
	//fmt.Println("format=", format)
	//format2 := now.Format("01-02")
	//fmt.Println("format2=", format2)

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
		Time, _ := time.Parse("2006-01-02 15:04:05", "2021-03-23 08:00:00")
		fmt.Println("Time1=", Time)
		//fmt.Println("Now=", time.Now())
		Time, _ = time.ParseInLocation("2006-01-02 15:04:05", "2021-03-23 08:00:00", time.Local)
		fmt.Println("Time2=", Time)
	*/

	/*
		timeSec := time.Now().Unix()
		fmt.Println("timeSec1=", timeSec)
		if timeSec > 0 {
			timeSec := time.Unix(timeSec, 0).Format("2006-01-02 15:04:05")
			fmt.Println("timeSec2=", timeSec)
		}
		fmt.Println("timeSec3=", timeSec)
	*/

	/*
		loc, _ := time.LoadLocation("Europe/Berlin")
		const longForm = "Jan 2, 2006 at 3:04pm (MST)"
		t, _ := time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
		fmt.Println(t)
	*/

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
