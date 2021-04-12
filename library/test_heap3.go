package main

import (
	"container/heap"
	"fmt"
)

type Item struct {
	value int // 对应id
	score int // 对应排序字段
}

// 最小堆
type MinHeap []*Item

func (mh MinHeap) Len() int { return len(mh) }
func (mh MinHeap) Less(i, j int) bool {
	return mh[i].score < mh[j].score
}
func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}
func (mh *MinHeap) Push(x interface{}) {
	item := x.(*Item)
	*mh = append(*mh, item)
}
func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	item := old[n-1]
	*mh = old[0 : n-1]
	return item
}

func Example_MinHeap() {
	mh := make(MinHeap, 0)
	items := map[int]int{
		111: 4, 222: 2, 333: 3, 444: 1,
	}
	for value, priority := range items {
		item := &Item{
			value: value,
			score: priority,
		}
		heap.Push(&mh, item)
		if mh.Len() > 2 {
			heap.Pop(&mh)
		}
	}
	for mh.Len() > 0 {
		item := heap.Pop(&mh).(*Item)
		fmt.Printf("%.2d:%d ", item.score, item.value)
	}
}

func main() {
	Example_MinHeap()
}
