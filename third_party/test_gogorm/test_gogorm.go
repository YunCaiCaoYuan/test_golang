package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type FacadeType int32

const (
	FacadeType_NONE   FacadeType = 0
	FacadeType_BRIGHT FacadeType = 1
	FacadeType_DARK   FacadeType = 2
)

type Player struct {
	Id       int64  `gorm:"column:Id"`
	Id2      int64  `gorm:"column:Id2"`
	Nickname string `gorm:"column:Nickname"`
	Sex      bool   `gorm:"column:Sex"`
	Charm    int64  `gorm:"column:Charm"`
	Wealth   int64  `gorm:"column:Wealth"`
	Flags    int64  `gorm:"column:Flags"`
	//Flags2      int64      `gorm:"column:Flags2"`
	//Flags3      int64      `gorm:"column:Flags3"`
	Icon        string     `gorm:"column:Icon"`
	OnlineExp   uint32     `gorm:"column:OnlineExp"`
	ClanId      int32      `gorm:"column:ClanId"`
	CreateAt    time.Time  `gorm:"column:CreateAt"`
	Point       int64      `gorm:"column:Point"`
	Signature   string     `gorm:"column:Signature"`
	City        string     `gorm:"column:City"`
	Birthday    string     `gorm:"column:Birthday"`
	OnlineState int        `gorm:"-"`
	IsDelete    bool       `gorm:"-"`
	Facade      FacadeType `gorm:"column:Facade"`
}

func (*Player) TableName() string {
	return "player"
}

func main() {
	dbDSN := "root:123456@tcp(127.0.0.1:3306)/test?parseTime=True"
	db, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	if err != nil {
		fmt.Println("err=", err)
	}

	res := db.Model(&Player{}).Where(" `Id`=? and `Sex`=0 ", 1001082).Updates(map[string]interface{}{
		"Sex":      1,
		"CreateAt": time.Now(),
	})
	if res.Error != nil {
		fmt.Println("res.Error=", res.Error)
	}
}
