package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
	"testing"
)

var (
	log *zap.Logger
	db  *gorm.DB
)

func start() {
	dbDSN := "root:123456@tcp(127.0.0.1:3306)/test?parseTime=True"
	db, _ = gorm.Open("mysql", dbDSN)
	var err error
	if err != nil {
		fmt.Println("err=", err)
	}
	log, _ = zap.NewProduction()
}

type testPlayer struct {
	Id  int64 `gorm:"column:id"`
	Id2 int64 `gorm:"column:Id2"`
}

// 要指定gorm tag
func Test_StructTag(t *testing.T) {
	start()
	list := make([]*testPlayer, 0)
	sql := "select id, Id2 from player where sex>0"
	if err := db.Raw(sql).Scan(&list).Error; err != nil {
		fmt.Println("err=", err)
		return
	}
	log.Info("Test_StructTag", zap.Any("list", list))
}
