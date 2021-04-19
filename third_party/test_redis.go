package main

import (
	"fmt"
	"log"
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

	const incrScript = `
local mt = setmetatable(_G, nil)
function serialize(obj)
    local lua = ""
    local t = type(obj)
    if t == "number" then
        lua = lua .. obj
    elseif t == "boolean" then
        lua = lua .. tostring(obj)
    elseif t == "string" then
        lua = lua .. string.format("%q", obj)
    elseif t == "table" then
        lua = lua .. "{\n"
        for k, v in pairs(obj) do
            lua = lua .. "[" .. serialize(k) .. "]=" .. serialize(v) .. ",\n"
        end
        local metatable = getmetatable(obj)
        if metatable ~= nil and type(metatable.__index) == "table" then
            for k, v in pairs(metatable.__index) do
                lua = lua .. "[" .. serialize(k) .. "]=" .. serialize(v) .. ",\n"
            end
        end
        lua = lua .. "}"
    elseif t == "nil" then
        return nil
    else
        error("can not serialize a " .. t .. " type.")
    end
    return lua
end
function unserialize(lua)
    local t = type(lua)
    if t == "nil" or lua == "" then
        return nil
    elseif t == "number" or t == "string" or t == "boolean" then
        lua = tostring(lua)
    else
        error("can not unserialize a " .. t .. " type.")
    end
    lua = "return " .. lua
    local func = loadstring(lua)
    if func == nil then
        return nil
    end
    return func()
end
setmetatable(_G, mt)

local dur_data_str = redis.call('GET', KEYS[1]);
local dur_data = unserialize(dur_data_str)
if not dur_data then
    dur_data = {}
    dur_data["updated_at"] = ARGV[1]
else
    local diff = ARGV[1] - dur_data["updated_at"]
    if 1 < diff and diff < 120 then
        if not dur_data["duration"] then
            dur_data["duration"] = 0
        end
        dur_data["duration"] = dur_data["duration"] + diff
    end
    dur_data["updated_at"] = ARGV[1]
end;
dur_data_str = serialize(dur_data)
redis.call('SET', KEYS[1], dur_data_str);
redis.call('EXPIRE', KEYS[1], 7*24*3600);
return 1
	`
	key := fmt.Sprintf("SC:PlayerDailyDuration:%d:%d:%s", 101, 102, "20210102")
	now := time.Now().Unix()
	res, err := redis.NewScript(incrScript).Run(RedisCon, []string{key}, now).Int()
	if err != nil {
		log.Println("err happen:", err)
	}
	log.Println("res:", res)

	/*
		const incrScript = `
			local num = redis.call('GET', KEYS[1]);
			if not num then
				redis.call('SET',KEYS[1], 1);
				return 1;
			else
				local res = num + 1;
				redis.call('SET',KEYS[1], res);
				return res;
			end
		`
		res, err := redis.NewScript(incrScript).Run(RedisCon, []string{fmt.Sprintf("%s%d", "KeySubChannelSeqNum", 111),}).Int()
	*/

	/*
	key := "v1.0:AccountSkipId"
	ids := []string{"1", "2"}
	err := RedisCon.SAdd(key, ids).Err()
	if err != nil {
		fmt.Println("AddSkipIds redis failed ", err)
	}*/

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
