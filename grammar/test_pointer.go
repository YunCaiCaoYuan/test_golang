package main

import (
	"fmt"
	"time"
)

type Nameplate struct {
	Id        int64     `json:"id" gorm:"Column:id;PRIMARY_KEY"`
	ChannelId int64     `json:"channel_id" gorm:"Column:channel_id"`
	Name      string    `json:"name" gorm:"Column:name"`
	Style     int       `json:"style" gorm:"Column:style"`
	CreatedAt time.Time `json:"created_at" gorm:"Column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"Column:updated_at"`
}

func func1() *Nameplate {
	var nameplate Nameplate
	fmt.Println("func1 nameplate=", nameplate)
	fmt.Println("func1 &nameplate=", &nameplate) // 有值
	//return &nameplate
	return nil
}

func func2() *Nameplate {
	return &Nameplate{}
}

func func3() Nameplate {
	ret := Nameplate{}
	fmt.Printf("ret10 = %p\n", &ret)
	return ret
}

func func4() *Nameplate {
	ret := Nameplate{}
	fmt.Printf("ret20 = %p\n", &ret)
	return &ret
}

func main() {
	res := func2()
	fmt.Println("res2=", res)

	ret := func3()
	fmt.Printf("ret11 = %p\n", &ret)

	ret1 := func4()
	fmt.Printf("ret22 = %p\n", ret1)  // 传指针vs传值，传值会发生深拷贝

	/*
	var res *Nameplate
	fmt.Println("res1=", res)

	res = func1()
	fmt.Println("res3=", res)
	 */

	/*
	nameplate := new(Nameplate)
	fmt.Println("nameplate1=", nameplate, "&nameplate", &nameplate)
	nameplate = func1()
	fmt.Println("nameplate2=", nameplate, "&nameplate", &nameplate)

	fmt.Println("sizeof &nameplate=", unsafe.Sizeof(nameplate))  // 8
	fmt.Println("sizeof *nameplate=", unsafe.Sizeof(*nameplate)) // 88

	fmt.Println("sizeof string=", unsafe.Sizeof(nameplate.Name)) // 16字节
	fmt.Println("sizeof time.Time=", unsafe.Sizeof(nameplate.CreatedAt)) // 24字节
	fmt.Println("sizeof empty struct", unsafe.Sizeof(struct {}{}))	// 0
	 */
}
