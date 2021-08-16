package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	redMax  = 33
	buleMax = 16
	redNum  = 6
	buleNum = 1
)
const Cnt = 5

func randNum(max, cnt int) sort.IntSlice {
	numMap := make(map[int]bool)
	for {
		num := rand.Intn(max) + 1
		if !numMap[num] {
			numMap[num] = true
		}
		if len(numMap) >= cnt {
			break
		}
	}
	list := make(sort.IntSlice, 0, max)
	for num, _ := range numMap {
		list = append(list, num)
	}
	list.Sort()
	return list
}

func wrap(num int) string {
	return fmt.Sprintf("%02d", num)
}

func main() {
	rand.Seed(time.Now().Unix())
	cnt := 0
	for {
		total := 0
		bigNum := 0
		list1 := randNum(redMax, redNum)
		for i := 0; i < len(list1); i++ {
			total += list1[i]
			if list1[i] >= 18 {
				bigNum++
			}
		}
		if total >= 88 && total <= 117 {
			// 去除大小15或06
			if bigNum <= 1 || bigNum >= 5 {
				continue
			}
			cnt += 1
			fmt.Println(wrap(list1[0]), wrap(list1[1]), wrap(list1[2]), wrap(list1[3]), wrap(list1[4]), wrap(list1[5]))
		}
		if cnt >= 16 {
			break
		}
	}
}
