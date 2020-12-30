package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type Player struct {
	Id          int64     `gorm:"column:Id"`
	Id2         int64     `gorm:"column:Id2"`
	Nickname    string    `gorm:"column:Nickname"`
	Sex         int8      `gorm:"column:Sex"`
	Charm       int64     `gorm:"column:Charm"`
	Wealth      int64     `gorm:"column:Wealth"`
	Flags       int64     `gorm:"column:Flags"`
	Icon        string    `gorm:"column:Icon"`
	OnlineTime  int64     `gorm:"column:OnlineTime"`
	OnlineExp   uint32    `gorm:"column:OnlineExp"`
	ClanId      int32     `gorm:"column:ClanId"`
	CreateAt    time.Time `gorm:"-"`
	IsDelete    bool      `gorm:"-"`
	CharmLevel  int32     `gorm:"-"`
	WealthLevel int32     `gorm:"-"`
}

func (_ *Player) TableName() string {
	return "player"
}

type ChannelRel struct {
	Id      int64 `json:"id,omitempty" gorm:"column:id"` // 自增 id
	OwnerId int64 `json:"owner_id,omitempty"`            // 创建者 id，只有主频道有
}

func (_ *ChannelRel) TableName() string {
	return "channel_rel"
}

func main() {
	dbDSN := "root:123456@tcp(139.159.144.149:3306)/test"
	db, err := gorm.Open("mysql", dbDSN)
	if err != nil {
		fmt.Println("err=", err)
	}

	/*
		// 提供TableName方法
		isFind := db.HasTable(&Player{})
		fmt.Println("isFind=", isFind)
		//ret := Player{}
		player := new(Player)
		err = db.Where("id = 1").First(player).Error
		if err != nil {
			fmt.Println("err=", err)
		}
		fmt.Println("ret=", player)
	*/

	// 推导的表名不对 channel_rels
	//id := 20000101
	//ret := new(ChannelRel)
	//err = db.Where("id = ?", id).First(ret).Error // Error 1146: Table 'test.channel_rels' doesn't exist
	//err = db.Where("id = ?", id).Find(ret).Error // Error 1146: Table 'test.channel_rels' doesn't exist
	//err = db.Table("channel_rel").Where("id = ?", id).Find(ret).Error
	//err = db.Table("channel_rel").Where("id = ?", id).First(ret).Error

	ret := make([]ChannelRel, 0)
	err = db.Where("owner_id = ?", 0).Order("created_at asc").Find(&ret).Error
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("ret=", ret)

	/*
		type countInfo struct {
			Count int32
		}
		var count countInfo
		err = db.Raw("select count(*) as count from player").Scan(&count).Error
		//count := 0
		//err = db.Raw("select count(*) from player").Count(&count).Error
		if err != nil {
			fmt.Println("err=", err)
		}
		fmt.Println("count=", count, "count.Count=", count.Count)
	*/

	/*
		players := make([]Player, 0)
		err = db.Raw("select * from player").Where(" id = ? ", 1).Scan(&players).Error // 语法错误
		if err != nil {
			fmt.Println("err=", err)
		}
		fmt.Println("players=", players)
	*/

	/*
		ret := Player{}
		err = db.Raw("select * from player where id = 123456").Scan(&ret).Error
		if err != nil {
			fmt.Println("err=", err) // record not found
		}
	*/
	/*
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
	*/

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
