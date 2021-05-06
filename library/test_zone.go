package main

import (
	"fmt"
	"time"
)

func main() {
	// golang 特定时间格式 2006-01-02 15:04:05
	// 按照美式时间格式非常容易记忆
	// 格式：月/日 小时：分钟：秒 年 外加时区
	// 01/02 03:04:05pm 06 -0700
	timeFormat := "2006-01-02 15:04:05"

	// 时区
	Loc, _ := time.LoadLocation("Asia/Shanghai")

	t := time.Now().In(Loc)
	fmt.Println(t.Unix()) //时间戳
	fmt.Println(t.Format(timeFormat)) //自定义格式

	// 日期转时间戳
	fmt.Println("datetime to timestamp")
	t2, _ := time.ParseInLocation(timeFormat, "2021-02-07 15:46:14", Loc)
	fmt.Println(t2.Unix())
	fmt.Println(t2.Format(timeFormat))

	// 时间戳转日期
	fmt.Println("timestamp to datetime")
	timestamp := int64(1612683974)
	t3 := time.Unix(timestamp, 0).In(Loc)
	fmt.Println(t3.Unix())
	fmt.Println(t3.Format(timeFormat))
}
