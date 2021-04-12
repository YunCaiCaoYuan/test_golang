package main

import "fmt"

type ChannelRel struct {
	ID int64 `json:"id,omitempty" gorm:"column:id"` // 自增 id
}

func main() {
	/*
	rels := make([]*ChannelRel, 0)
	fmt.Println(len(rels))
	fmt.Println(rels)*/

	slice := []int{1,2}
	fmt.Println(slice[:4]) // panic: runtime error: slice bounds out of range
}
