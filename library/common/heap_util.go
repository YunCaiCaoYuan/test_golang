package common

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
