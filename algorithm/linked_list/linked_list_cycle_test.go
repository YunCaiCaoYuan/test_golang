package linked_list

import (
	"fmt"
	"testing"
)

// hash表
func hasCycle(head *Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	visitMap := make(map[*Node]struct{})
	for head != nil {
		if _, ok := visitMap[head]; ok {
			return true
		}
		visitMap[head] = struct{}{}
		head = head.Next
	}
	return false
}

// 3->2->0->-4 pos=1
func Test_hasCycle1(t *testing.T) {
	node4 := &Node{Val: -4}
	node3 := &Node{Val: 0}
	node2 := &Node{Val: 2}
	node1 := &Node{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2
	fmt.Println(hasCycle(node1))
}

// 1->2 pos=0
func Test_hasCycle2(t *testing.T) {
	node2 := &Node{Val: 2}
	node1 := &Node{Val: 1}
	node1.Next = node2
	node2.Next = node1
	fmt.Println(hasCycle(node1))
}

// 1 pos=-1
func Test_hasCycle3(t *testing.T) {
	node1 := &Node{Val: 1}
	fmt.Println(hasCycle(node1))
}

// eg: go test -run=hasCycle
