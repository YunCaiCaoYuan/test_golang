package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type FacadeType int32

const (
	FacadeType_NONE   FacadeType = 0
	FacadeType_BRIGHT FacadeType = 1
	FacadeType_DARK   FacadeType = 2
)

type SexType int32

const (
	SexType_UNKNOWN SexType = 0
	SexType_MALE    SexType = 1
	SexType_FEMALE  SexType = 2
)

type Player struct {
	Id          int64      `gorm:"column:Id"`
	Id2         int64      `gorm:"column:Id2"`
	Nickname    string     `gorm:"column:Nickname"`
	Sex         SexType    `gorm:"column:Sex"`
	Charm       int64      `gorm:"column:Charm"`
	Wealth      int64      `gorm:"column:Wealth"`
	Flags       int64      `gorm:"column:Flags"`
	Flags2      int64      `gorm:"column:Flags2"`
	Flags3      int64      `gorm:"column:Flags3"`
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

func (_ *Player) TableName() string {
	return "player"
}

type ChannelRel struct {
	Id          int64     `json:"id,omitempty"`            // 频道id
	ParentId    int64     `json:"parent_id,omitempty"`     // 父亲频道 id, 主频道该项为 0
	ParentSeq   int32     `json:"parent_seq,omitempty"`    // 父亲频道 seq,主频道该项为 0
	TopParentId int64     `json:"top_parent_id,omitempty"` // 顶级父亲id
	Level       int32     `json:"level,omitempty"`         // 频道层级，0代表主频道
	Seq         int32     `json:"seq,omitempty"`           // 子频道序号
	Sort        int32     `json:"sort,omitempty"`          // 排序
	OwnerId     int64     `json:"owner_id,omitempty"`      // 创建者 id，只有主频道有
	ChildNum    int32     `json:"child_num,omitempty"`     // 下一级子频道的数量
	IsDefault   bool      `json:"is_default"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func main() {
	//dbDSN := "root:123456@tcp(139.159.144.149:3306)/test?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local"
	dbDSN := "root:123456@tcp(139.159.144.149:3306)/test?parseTime=True"
	db, err := gorm.Open("mysql", dbDSN)
	if err != nil {
		fmt.Println("err=", err)
	}

	// 提供TableName方法
	isFind := db.HasTable(&Player{})
	fmt.Println("isFind=", isFind)
	//ret := Player{}
	player := new(Player)
	err = db.Where("id = 999").First(player).Error
	if err != nil {
		fmt.Println("err=", err)
	}
	fmt.Println("player.Facade=", player.Facade)

	// 推导的表名不对 channel_rels
	//id := 20000101
	//ret := new(ChannelRel)
	//err = db.Where("id = ?", id).First(ret).Error // Error 1146: Table 'test.channel_rels' doesn't exist
	//err = db.Where("id = ?", id).Find(ret).Error // Error 1146: Table 'test.channel_rels' doesn't exist
	//err = db.Table("channel_rel").Where("id = ?", id).Find(ret).Error
	//err = db.Table("channel_rel").Where("id = ?", id).First(ret).Error

	//ret := make([]ChannelRel, 0)
	//err = db.Where("owner_id = ?", 0).Order("created_at asc").Find(&ret).Error
	//if err != nil {
	//	fmt.Println("err=", err)
	//}
	//fmt.Println("ret=", ret)

	/*
		count := 0
		err = db.Raw("select count(*) from player").Count(&count).Error
		if err != nil {
			fmt.Println("err=", err)
		}
		fmt.Println("count=", count)
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
