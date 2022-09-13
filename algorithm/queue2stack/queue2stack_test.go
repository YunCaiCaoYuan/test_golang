package queue2stack

import (
	"fmt"
	"testing"
)

// queue 操作
func initQueue(size int64) []int64 {
	return make([]int64, 0, size)
}
func pushback(q []int64, val int64) []int64 {
	return append(q, val)
}
func peekQueue(q []int64) int64 {
	return q[0]
}
func popQueue(q []int64) (int64, []int64) {
	return q[0], q[1:len(q)]
}
func isQueueEmpty(q []int64) bool {
	return len(q) == 0
}

type SpecialStack struct {
	LeftQueue  []int64
	RightQueue []int64
}

// stack impl
func NewStack(size int64) *SpecialStack {
	return &SpecialStack{
		LeftQueue:  initQueue(size),
		RightQueue: initQueue(size),
	}
}
func Push(s *SpecialStack, val int64) *SpecialStack {
	if !isQueueEmpty(s.RightQueue) {
		return &SpecialStack{
			s.LeftQueue,
			pushback(s.RightQueue, val),
		}
	}
	return &SpecialStack{
		pushback(s.LeftQueue, val),
		s.RightQueue,
	}
}
func Peek(s *SpecialStack) int64 {
	if isQueueEmpty(s.RightQueue) {
		return s.LeftQueue[len(s.LeftQueue)-1]
	}
	return s.RightQueue[len(s.RightQueue)-1]
}
func Pop(s *SpecialStack) (*SpecialStack, int64) {
	if isQueueEmpty(s.RightQueue) {
		return &SpecialStack{
			s.LeftQueue[:0],
			append(s.RightQueue, s.LeftQueue[0:len(s.LeftQueue)-1]...),
		}, s.LeftQueue[len(s.LeftQueue)-1]
	}
	return &SpecialStack{
		append(s.LeftQueue, s.RightQueue[0:len(s.RightQueue)-1]...),
		s.RightQueue[:0],
	}, s.RightQueue[len(s.RightQueue)-1]
}
func Empty(s *SpecialStack) bool {
	return isQueueEmpty(s.RightQueue) && isQueueEmpty(s.LeftQueue)
}

func Test_Queue2Stack(t *testing.T) {
	stack := NewStack(10)
	stack = Push(stack, 1)
	stack = Push(stack, 2)
	stack = Push(stack, 3)
	stack, r := Pop(stack)
	fmt.Println("r:", r)
	stack, r = Pop(stack)
	fmt.Println("r:", r)
	stack = Push(stack, 4)
	stack = Push(stack, 5)
	stack, r = Pop(stack)
	fmt.Println("r:", r)
	stack, r = Pop(stack)
	fmt.Println("r:", r)
	stack, r = Pop(stack)
	fmt.Println("r:", r)
}

// 该方法不对，没有使用队列的方法
