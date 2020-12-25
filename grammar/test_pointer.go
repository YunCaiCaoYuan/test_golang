package main

import (
	"fmt"
	"time"
	"unsafe"
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

func main() {
	nameplate := new(Nameplate)
	fmt.Println("nameplate1=", nameplate, "&nameplate", &nameplate)
	nameplate = func1()
	fmt.Println("nameplate2=", nameplate, "&nameplate", &nameplate)

	fmt.Println("sizeof &nameplate=", unsafe.Sizeof(nameplate))
	fmt.Println("sizeof *nameplate=", unsafe.Sizeof(*nameplate))

	fmt.Println("sizeof string=", unsafe.Sizeof(nameplate.Name)) // 16字节
}
