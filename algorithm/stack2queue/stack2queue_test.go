package stack2queue

import (
	"fmt"
	"testing"
)

type Queue struct {
	inStack  []int64
	outStack []int64
}

func NewQueue() *Queue {
	return &Queue{
		inStack:  make([]int64, 0),
		outStack: make([]int64, 0),
	}
}

func (q *Queue) In(v int64) {
	q.inStack = append(q.inStack, v)
}
func (q *Queue) Out() int64 {
	if len(q.outStack) == 0 {
		for i := len(q.inStack) - 1; i >= 0; i-- {
			q.outStack = append(q.outStack, q.inStack[i])
		}
		q.inStack = q.inStack[:0]
	}
	ret := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return ret
}
func (q *Queue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}

func Test_Queue(t *testing.T) {
	q := NewQueue()
	q.In(1)
	q.In(2)
	q.In(3)
	fmt.Println(q.Out())
	fmt.Println(q.Out())
	q.In(4)
	q.In(5)
	fmt.Println(q.Out())
	fmt.Println(q.Out())
	fmt.Println(q.Out())
}

// 1 2 3 4 5
