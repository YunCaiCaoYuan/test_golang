package main

import (
	"fmt"
	//"git.2tianxin.com/platform/utils/uredis"
	"github.com/go-redis/redis"
	"strings"
	"time"
)

func main() {
	host := "127.0.0.1:6379"
	pwd := ""
	db := 0
	RedisCon := New(host, pwd, db)

	key := "v1.0:AccountSkipId"
	ids := []string{"1", "2"}
	err := RedisCon.SAdd(key, ids).Err()
	if err != nil {
		fmt.Println("AddSkipIds redis failed ", err)
	}

	/*
		fmt.Println("time.Minute=", time.Minute)
		fmt.Println("20000376 % 8 =", 20000376 % 8)
		now := GetNowUnix()
		fmt.Println("now=", now)
		ret := strconv.FormatInt(GetNowUnix(), 10)
		fmt.Println("ret=", ret)
		now2 := time.Now().Unix()
		fmt.Println("now2=", now2)
	*/
}

func New(host string, pwd string, db int) *redis.Client {
	configs := strings.Split(host, ",")
	con := redis.NewClient(&redis.Options{
		PoolSize: 100,
		Addr:     strings.TrimSpace(configs[0]),
		Password: pwd,
		DB:       db,
	})
	_, err := con.Ping().Result()
	if err != nil {
		panic("redis初始化失败, err:" + err.Error())
	}
	return con
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
