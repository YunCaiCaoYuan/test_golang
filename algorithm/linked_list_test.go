package algorithm

import (
	"testing"
)

type node struct {
	val  int
	next *node
}

func reverseList(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}

	p1 := head
	p2 := head.next
	for p2 != nil {
		p3 := p2.next
		p2.next = p1
		p1 = p2
		p2 = p3
	}
	head.next = nil
	head = p1

	return head
}

func printList(head *node) {
	for head != nil {
		print(head.val, " ")
		head = head.next
	}
	print("\n")
}

// 1->2->3->4->5
// 5->4->3->2->1
func Test_reverseList1(t *testing.T) {
	nod := &node{val: 1}
	tmp := nod
	for i := 1; i < 5; i++ {
		newNod := &node{val: i + 1}
		tmp.next = newNod
		tmp = newNod
	}
	printList(nod)

	print("after reverse:\n")
	printList(reverseList(nod))
}

// 1
func Test_reverseList2(t *testing.T) {
	nod := &node{val: 1}
	printList(nod)

	print("after reverse:\n")
	printList(reverseList(nod))
}

// nil
func Test_reverseList3(t *testing.T) {
	printList(nil)

	print("after reverse:\n")
	printList(reverseList(nil))
}
