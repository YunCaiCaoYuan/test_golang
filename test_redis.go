package main

import (
	"fmt"
	"strconv"
	"time"
	//"github.com/go-redis/redis"
)

func main() {
	//fmt.Println("time.Minute=", time.Minute)

	//fmt.Println("20000376 % 8 =", 20000376 % 8)

	now := GetNowUnix()
	fmt.Println("now=", now)
	ret := strconv.FormatInt(GetNowUnix(), 10)
	fmt.Println("ret=", ret)
	now2 := time.Now().Unix()
	fmt.Println("now2=", now2)
}

//获取当前时间戳
func GetNowUnix() int64 {
	return GetNow().Unix()
}

//获取当前时间
var OffsetTime int64 = 0

func GetNow() time.Time {
	now := time.Now().Add(time.Duration(OffsetTime) * time.Second)
	return now
}
