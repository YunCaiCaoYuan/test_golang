package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	dbDSN := "root:123456@tcp(139.159.144.149:3306)/test"
	db, err := gorm.Open("mysql", dbDSN)
	if err != nil {
		fmt.Println("err=", err)
	}

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
