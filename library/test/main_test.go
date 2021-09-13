package test

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var db struct {
	Dns string
}

func TestMain(m *testing.M) {
	db.Dns = os.Getenv("DATABASE_DNS")
	if db.Dns == "" {
		db.Dns = "root:123456@tcp(localhost:3306)/?charset=utf8&parseTime=True&loc=Local"
	}

	flag.Parse()
	exitCode := m.Run()

	db.Dns = ""

	// 退出
	os.Exit(exitCode)
}

func TestDatabase(t *testing.T) {
	fmt.Println(db.Dns)
}
