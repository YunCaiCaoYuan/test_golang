package test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

type MsgData struct {
	X, Y, Z int
	Name    string
}

var network bytes.Buffer //网络传递的数据载体

func Test_Gob(t *testing.T) {
	err := senMsg()
	if err != nil {
		fmt.Println("编码错误")
		return
	}
	err = revMsg()
	if err != nil {
		fmt.Println("解码错误")
		return
	}
}

func senMsg() error {
	fmt.Print("开始执行编码（发送端）")

	enc := gob.NewEncoder(&network)
	sendMsg := MsgData{3, 4, 5, "jiangzhou"}
	fmt.Println("原始数据：", sendMsg)
	err := enc.Encode(&sendMsg)
	fmt.Println("传递的编码数据为：", network)
	return err
}
func revMsg() error {
	var revData MsgData
	dec := gob.NewDecoder(&network)
	err := dec.Decode(&revData) //传递参数必须为 地址
	fmt.Println("解码之后的数据为：", revData)
	return err
}
