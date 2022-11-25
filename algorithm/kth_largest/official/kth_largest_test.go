package official

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"
)

type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, val := range nums {
		kl.Add(val)
	}
	return kl
}

func (kl *KthLargest) Push(v interface{}) {
	kl.IntSlice = append(kl.IntSlice, v.(int))
}

func (kl *KthLargest) Pop() interface{} {
	a := kl.IntSlice
	v := a[len(a)-1]
	kl.IntSlice = a[:len(a)-1]
	return v
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(kl, val)
	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
	return kl.IntSlice[0]
}

func Test_1(t *testing.T) {
	kl := Constructor(3, []int{6, 1, 91})
	fmt.Println("kl.IntSlice[0]", kl.IntSlice[0])
	fmt.Println("kl.Add(2)", kl.Add(2))
	fmt.Println("kl.Add(88)", kl.Add(88))
}

//kl.IntSlice[0] 1
//kl.Add(2) 2
//kl.Add(2) 6
