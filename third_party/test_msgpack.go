package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	"time"
)

type Player struct {
	Id          int64         `gorm:"column:Id"`
	Id2         int64         `gorm:"column:Id2"`
	Nickname    string        `gorm:"column:Nickname"`
	Sex         int32    `gorm:"column:Sex"`
	Charm       int64         `gorm:"column:Charm"`
	Wealth      int64         `gorm:"column:Wealth"`
	Flags       int64         `gorm:"column:Flags"`
	Flags2      int64         `gorm:"column:Flags2"`
	Flags3      int64         `gorm:"column:Flags3"`
	Icon        string        `gorm:"column:Icon"`
	OnlineExp   uint32        `gorm:"column:OnlineExp"`
	ClanId      int32         `gorm:"column:ClanId"`
	CreateAt    time.Time     `gorm:"column:CreateAt"`
	Point       int64         `gorm:"column:Point"`
	Signature   string        `gorm:"column:Signature"`
	City        string        `gorm:"column:City"`
	Birthday    string        `gorm:"column:Birthday"`
	OnlineState int           `gorm:"-"`
	IsDelete    bool          `gorm:"-"`
	Facade      int32 `gorm:"column:Facade"`
}

func main() {
	buf := []byte{-17, -65, -67, 0, 20, -17, -65, -67, 73, 100, -17, -65, -67, 0, 0, 0, 0, 0, 15, 66, 85, -17, -65, -67, 73, 100, 50, -17, -65, -67, 0, 0, 0, 0, 0, 0, 0, 0, -17, -65, -67, 78, 105, 99, 107, 110, 97, 109, 101, -17, -65, -67, -26, -106, -80, -26, -119, -117, 49, 48, 48, 48, 48, 50, 49, -17, -65, -67, 83, 101, 120, -17, -65, -67, 0, 0, 0, 1, -17, -65, -67, 67, 104, 97, 114, 109, -17, -65, -67, 0, 0, 0, 0, 0, -44, -128, 40, -17, -65, -67, 87, 101, 97, 108, 116, 104, -17, -65, -67, 0, 0, 0, 0, 0, 0, 0, 0, -17, -65, -67, 70, 108, 97, 103, 115, -17, -65, -67, 0, 0, 0, 0, 0, 0, 0, 0, -17, -65, -67, 70, 108, 97, 103, 115, 50, -17, -65, -67, 0, 0, 0, 0, 0, 0, 0, 0, -17, -65, -67, 70, 108, 97, 103, 115, 51, -17, -65, -67, 0, 0, 0, 0, 0, 0, 0, 0, -17, -65, -67, 73, 99, 111, 110, -17, -65, -67, -17, -65, -67, 79, 110, 108, 105, 110, 101, 69, 120, 112, -17, -65, -67, 0, 0, 0, -17, -65, -67, -17, -65, -67, 67, 108, 97, 110, 73, 100, -17, -65, -67, 0, 0, 0, 0, -17, -65, -67, 67, 114, 101, 97, 116, 101, 65, 116, -17, -65, -67, -17, -65, -67, 92, -37, -117, -62, -91, 80, 111, 105, 110, 116, -17, -65, -67, 0, 0, 0, 0, 0, 0, 0, 0, -17, -65, -67, 83, 105, 103, 110, 97, 116, 117, 114, 101, -17, -65, -67, -17, -65, -67, 67, 105, 116, 121, -17, -65, -67, -17, -65, -67, 66, 105, 114, 116, 104, 100, 97, 121, -17, -65, -67, -17, -65, -67, 79, 110, 108, 105, 110, 101, 83, 116, 97, 116, 101, -17, -65, -67, 0, 0, 0, 0, 0, 0, 0, 0, -17, -65, -67, 73, 115, 68, 101, 108, 101, 116, 101, -62, -90, 70, 97, 99, 97, 100, 101, -17, -65, -67, 0, 0, 0, 0}
	player := new(Player)
	err := msgpack.Unmarshal(buf, player)
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("player=", player)

	// constant -17 overflows byte
}



