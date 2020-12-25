package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

func main() {
	dbDSN := "root:123456@tcp(139.159.144.149:3306)/test"
	db, err := gorm.Open("mysql", dbDSN)
	if err != nil {
		fmt.Println("err=", err)
	}
	type Player struct {
		Id          int64      `gorm:"column:Id"`
		Id2         int64      `gorm:"column:Id2"`
		Nickname    string     `gorm:"column:Nickname"`
		Sex         int8 	    `gorm:"column:Sex"`
		Charm       int64      `gorm:"column:Charm"`
		Wealth      int64      `gorm:"column:Wealth"`
		Flags       int64      `gorm:"column:Flags"`
		Icon        string     `gorm:"column:Icon"`
		OnlineTime  int64      `gorm:"column:OnlineTime"`
		OnlineExp   uint32     `gorm:"column:OnlineExp"`
		ClanId      int32      `gorm:"column:ClanId"`
		CreateAt    time.Time  `gorm:"column:CreateAt"`
		IsDelete    bool       `gorm:"-"`
		CharmLevel  int32      `gorm:"-"`
		WealthLevel int32      `gorm:"-"`
	}
	players := make([]*Player, 0)
	ids := []int32{1,2,6}
	playerIds := make([]int64, 0)
	err = db.Raw("select Id FROM player WHERE id in (?) ", ids).Scan(&players).Error
	if err != nil {
		fmt.Println("err=", err)
	}
	for _, item := range players {
		fmt.Println("item=", item)
		playerIds = append(playerIds, item.Id)
	}
	fmt.Println("playerIds=", playerIds)

	/*
	type TbUser struct {
		Id       int32
		Username string
	}
	tbUsers := make([]*TbUser, 0)
	err = db.Raw("select * from tb_user").Scan(&tbUsers).Error
	if err != nil {
		fmt.Println("err", err)
	}
	for _, tbUser := range tbUsers {
		fmt.Println("tbUser=", tbUser)
	}
	 */

	/*
		type Product struct {
			gorm.Model
			Code  string
			Price uint
		}
		db, err := gorm.Open("sqlite3", "test.db")
		if err != nil {
			panic("failed to connect database")
		}
		defer db.Close()

		// Migrate the schema
		db.AutoMigrate(&Product{})

		// 创建
		db.Create(&Product{Code: "L1212", Price: 1000})

		// 读取
		var product Product
		db.First(&product, 1)                   // 查询id为1的product
		db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

		// 更新 - 更新product的price为2000
		db.Model(&product).Update("Price", 2000)

		// 删除 - 删除product
		db.Delete(&product)
	*/
}
