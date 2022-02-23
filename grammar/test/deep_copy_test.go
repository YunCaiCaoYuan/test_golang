package test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"
)

//速度速值
type Speed int

//风扇转速
type FanSpeed struct {
	Speed Speed
}

//售价
type Money struct {
	Length float64
}

//内存数量以及大小
type Memory struct {
	Count      int
	MemorySize []int
}

//电脑信息
type Computer struct {
	SystemName string              //系统名字
	UseNumber  int                 //使用次数
	Memory     Memory              //存储
	Fan        map[string]FanSpeed //风扇
	Money      Money               //售价
}

func ModifyCat1(pc Computer) {
	fmt.Printf("PcInfo Pc1:%v\n", pc)
	pc.SystemName = "MacOs"
	pc.UseNumber = 100
	pc.Memory.Count = 2
	pc.Memory.MemorySize[0] = 8
	pc.Memory.MemorySize[1] = 8
	pc.Memory.MemorySize[2] = 0
	pc.Memory.MemorySize[3] = 0
	pc.Fan["left"] = FanSpeed{2000}
	pc.Fan["right"] = FanSpeed{1500}
	fmt.Printf("PcInfo Pc1:%v\n", pc)
}

func deepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func Test_ComputerStart4(t *testing.T) {
	Pc1 := &Computer{
		SystemName: "Windows",
		UseNumber:  1000,
		Memory:     Memory{Count: 4, MemorySize: []int{32, 32, 32, 32}},
		Fan:        map[string]FanSpeed{"left": {2500}, "right": {2000}},
		Money:      Money{123.45},
	}

	//深拷贝
	Pc2 := new(Computer)
	if err := deepCopy(Pc2, Pc1); err != nil {
		panic(err.Error())
	}
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", Pc1, Pc2)

	ModifyCat1(*Pc2)
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", Pc1, Pc2)
}
