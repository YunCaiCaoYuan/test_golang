package test

import (
	"container/list"
	"fmt"
	"testing"
)

func Test_container_list(t *testing.T) {
	l := list.New()
	//fmt.Println("l0:", l)
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	//fmt.Println("l1:", l)

	ne := NewElementWrapImpl(l)
	ne.accept(&MyVisitorImpl{})
}

// 访问者模式
type MyVisitor interface {
	Visit(l *list.List)
}

type MyVisitorImpl struct{}

func (this *MyVisitorImpl) Visit(l *list.List) {
	if l.Len() > 0 {
		first := l.Front()
		for {
			fmt.Printf("%v ", first.Value)
			if first.Next() == nil {
				fmt.Println("")
				break
			}

			first = first.Next()
		}
	}
}

type ElementWrap interface {
	accept(v MyVisitor)
}

func NewElementWrapImpl(l *list.List) *ElementWrapImpl {
	return &ElementWrapImpl{L: l}
}

type ElementWrapImpl struct {
	L *list.List
}

func (this *ElementWrapImpl) accept(v MyVisitor) {
	v.Visit(this.L)
}
