package main

// 解析json配置

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Host string `json:"host"`
	Port int32  `json:"port"`
}

func main() {
	bin, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	config := new(Config)
	err = json.Unmarshal(bin, &config)
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	fmt.Println("config=", config)
}

// config= &{localhost 27017}
