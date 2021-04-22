package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Item struct {
	Value int // 对应id
	Score int // 对应排序字段
}

// 最小堆
type MinHeap []*Item

func (mh MinHeap) Len() int { return len(mh) }
func (mh MinHeap) Less(i, j int) bool {
	return mh[i].Score < mh[j].Score
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
		111: 3, 222: 2, 333: 4, 444: 1, 555: 9, 666 :8,
	}
	for value, priority := range items {
		item := &Item{
			Value: value,
			Score: priority,
		}
		heap.Push(&mh, item)
		if mh.Len() > 4 {
			heap.Pop(&mh)
		}
	}
	//for mh.Len() > 0 {
	//	item := heap.Pop(&mh).(*Item)
	//	fmt.Printf("%.2d:%d ", item.Score, item.Value)
	//}
	// 升序
	sort.Slice(mh, func(i, j int) bool {
		return mh[i].Score > mh[j].Score
	})
	for _, item := range mh {
		fmt.Println(item.Value, " : ", item.Score)
	}
}

func main() {
	Example_MinHeap()
}