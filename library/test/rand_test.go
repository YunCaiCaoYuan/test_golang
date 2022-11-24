package test

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

// 带权重随机

type weightCount struct {
	weight int64
	bid    int64
}

func randWeight(list []*weightCount) int64 {
	sort.Slice(list, func(i, j int) bool {
		return list[i].bid < list[j].bid
	})
	var weightTotal int64
	for _, item := range list {
		weight := item.weight
		item.weight += weightTotal
		weightTotal += weight
	}
	//printArr(list)
	randW := rand.Int63n(weightTotal) + 1
	for _, item := range list {
		if randW <= item.weight {
			return item.bid
		}
	}
	return list[0].bid
}

func printArr(list []*weightCount) {
	for _, item := range list {
		fmt.Printf("%v-%v ", item.bid, item.weight)
	}
	fmt.Println("")
}

func Test_1(t *testing.T) {
	rand.Seed(time.Now().Unix())
	bidMap := make(map[int64]int64)
	for i := 1; i <= 10000; i++ {
		list := []*weightCount{{60, 1}, {10, 2}, {10, 3}, {10, 4}, {5, 5}, {5, 6}}
		bid := randWeight(list)
		bidMap[bid] += 1
	}
	fmt.Println(bidMap)
}

func Test_2(t *testing.T) {
	list := []*weightCount{{60, 1}}
	fmt.Println(randWeight(list))
}

func Test_3(t *testing.T) {
	list := []*weightCount{}
	fmt.Println(randWeight(list))
}
