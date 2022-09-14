package queue2stack

import (
	"fmt"
	"testing"
)

type StackV2 struct {
	mainQueue []int64
	tmpQueue  []int64
}

func NewStackV2() *StackV2 {
	return &StackV2{
		mainQueue: make([]int64, 0),
		tmpQueue:  make([]int64, 0),
	}
}

func (s *StackV2) Push(v int64) {
	s.tmpQueue = append(s.tmpQueue, v)
	s.tmpQueue = append(s.tmpQueue, s.mainQueue...)
	tmp := s.mainQueue
	s.mainQueue = s.tmpQueue
	s.tmpQueue = tmp[:0]
}
func (s *StackV2) Pop() int64 {
	top := s.mainQueue[0]
	s.mainQueue = s.mainQueue[1:len(s.mainQueue)]
	return top
}
func (s *StackV2) Empty() bool {
	return len(s.mainQueue) == 0
}

func Test_StackV2(t *testing.T) {
	s := NewStackV2()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	s.Push(4)
	s.Push(5)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}

// 3 2 5 4 1
