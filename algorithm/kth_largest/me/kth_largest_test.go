package me

import (
	"container/heap"
	"fmt"
	"testing"
)

type KthLargest struct {
	elem []int
	k    int
}

func (this *KthLargest) Len() int           { return len(this.elem) }
func (this *KthLargest) Less(i, j int) bool { return this.elem[i] < this.elem[j] }
func (this *KthLargest) Swap(i, j int)      { this.elem[i], this.elem[j] = this.elem[j], this.elem[i] }
func (this *KthLargest) Push(x interface{}) {
	this.elem = append(this.elem, x.(int))
}
func (this *KthLargest) Pop() interface{} {
	old := this.elem
	n := len(old)
	x := old[n-1]
	this.elem = old[0 : n-1]
	return x
}

func Constructor(k int, nums []int) KthLargest {
	obj := KthLargest{
		elem: make([]int, 0),
		k:    k,
	}
	for _, item := range nums {
		obj.Add(item)
	}
	return obj
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}
	return this.elem[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func Test_run(t *testing.T) {
	k := 3
	nums := []int{34, 1, 9}
	obj := Constructor(k, nums)
	param_1 := obj.Add(10)
	fmt.Println(param_1)
}

func Test_2(t *testing.T) {
	kl := Constructor(3, []int{6, 1, 91})
	fmt.Println("kl.IntSlice[0]", kl.elem[0])
	fmt.Println("kl.Add(2)", kl.Add(2))
	fmt.Println("kl.Add(88)", kl.Add(88))
}

//kl.IntSlice[0] 1
//kl.Add(2) 2
//kl.Add(88) 6
