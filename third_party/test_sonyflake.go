package main

import (
	"fmt"
	"github.com/sony/sonyflake"
	"time"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	st.StartTime = time.Now()
	sf = sonyflake.NewSonyflake(st)
}

func main() {
	id, err := sf.NextID()
	fmt.Println("id=", id)
	fmt.Println(err)
}
