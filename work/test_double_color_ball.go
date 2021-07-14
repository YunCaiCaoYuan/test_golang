package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	redMax = 33
	buleMax = 16
	redNum = 6
	buleNum = 1
)
const Cnt = 5

func randNum(max, cnt int) sort.IntSlice {
	numMap := make(map[int]bool)
	for {
		num := rand.Intn(max)+1
		if !numMap[num] {
			numMap[num] = true
		}
		if len(numMap) >= cnt {
			break
		}
	}
	list := make(sort.IntSlice, 0, max)
	for num,_ := range numMap {
		list = append(list, num)
	}
	list.Sort()
	return list
}

func main() {
	rand.Seed(time.Now().Unix())
	for i:=0; i<Cnt; i++ {
		list1 := randNum(redMax, redNum)
		list2 := randNum(buleMax, buleNum)
		fmt.Println(list1[0], list1[1], list1[2], list1[3], list1[4], list1[5], list2[0])
	}
}
