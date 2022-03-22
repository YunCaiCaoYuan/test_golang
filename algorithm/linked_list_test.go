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

// ------------------------------------------------------------
func swapList(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}

	p1 := head
	p2 := head.next
	for p1 != nil && p2 != nil {
		p1.val, p2.val = p2.val, p1.val
		p1 = p2.next
		if p1 != nil {
			p2 = p1.next
		}
	}
	return head
}

// 1->2
// 2->1
func Test_swapList1(t *testing.T) {
	nod := &node{val: 1}
	tmp := nod
	for i := 1; i < 2; i++ {
		newNod := &node{val: i + 1}
		tmp.next = newNod
		tmp = newNod
	}
	printList(nod)

	print("after swap:\n")
	printList(swapList(nod))
}

// 1->2->3
// 2->1->3
func Test_swapList2(t *testing.T) {
	nod := &node{val: 1}
	tmp := nod
	for i := 1; i < 3; i++ {
		newNod := &node{val: i + 1}
		tmp.next = newNod
		tmp = newNod
	}
	printList(nod)

	print("after swap:\n")
	printList(swapList(nod))
}

// 1
// 1
func Test_swapList3(t *testing.T) {
	nod := &node{val: 1}
	printList(nod)

	print("after swap:\n")
	printList(swapList(nod))
}

// nil
// nil
func Test_swapList4(t *testing.T) {
	var nod *node
	printList(nod)

	print("after swap:\n")
	printList(swapList(nod))
}
func Test_swapList5(t *testing.T) {
	nod := &node{val: 1}
	tmp := nod
	for i := 1; i < 11; i++ {
		newNod := &node{val: i + 1}
		tmp.next = newNod
		tmp = newNod
	}
	printList(nod)

	print("after swap:\n")
	printList(swapList(nod))
}

// eg: go test -v linked_list_test.go -run=swapList
