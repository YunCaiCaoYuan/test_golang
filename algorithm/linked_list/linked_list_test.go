package linked_list

import (
	"testing"
)

func reverseList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	p1 := head
	p2 := head.Next
	for p2 != nil {
		p3 := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = p3
	}
	head.Next = nil
	head = p1

	return head
}

func printList(head *Node) {
	for head != nil {
		print(head.Val, " ")
		head = head.Next
	}
}

// 1->2->3->4->5
// 5->4->3->2->1
func Test_reverseList1(t *testing.T) {
	nod := &Node{Val: 1}
	tmp := nod
	for i := 1; i < 5; i++ {
		newNod := &Node{Val: i + 1}
		tmp.Next = newNod
		tmp = newNod
	}
	printList(nod)

	print("after reverse:\n")
	printList(reverseList(nod))
}

// 1
func Test_reverseList2(t *testing.T) {
	nod := &Node{Val: 1}
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
func swapList(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	p1 := head
	p2 := head.Next
	for p1 != nil && p2 != nil {
		p1.Val, p2.Val = p2.Val, p1.Val
		p1 = p2.Next
		if p1 != nil {
			p2 = p1.Next
		}
	}
	return head
}

// 1->2
// 2->1
func Test_swapList1(t *testing.T) {
	nod := &Node{Val: 1}
	tmp := nod
	for i := 1; i < 2; i++ {
		newNod := &Node{Val: i + 1}
		tmp.Next = newNod
		tmp = newNod
	}
	printList(nod)

	print("after swap:\n")
	printList(swapList(nod))
}

// 1->2->3
// 2->1->3
func Test_swapList2(t *testing.T) {
	nod := &Node{Val: 1}
	tmp := nod
	for i := 1; i < 3; i++ {
		newNod := &Node{Val: i + 1}
		tmp.Next = newNod
		tmp = newNod
	}
	printList(nod)

	print("after swap:\n")
	printList(swapList(nod))
}

// 1
// 1
func Test_swapList3(t *testing.T) {
	nod := &Node{Val: 1}
	printList(nod)

	print("after swap:\n")
	printList(swapList(nod))
}

// nil
// nil
func Test_swapList4(t *testing.T) {
	var nod *Node
	printList(nod)

	print("after swap:\n")
	printList(swapList(nod))
}
func Test_swapList5(t *testing.T) {
	nod := &Node{Val: 1}
	tmp := nod
	for i := 1; i < 11; i++ {
		newNod := &Node{Val: i + 1}
		tmp.Next = newNod
		tmp = newNod
	}
	printList(nod)

	print("after swap:\n")
	printList(swapList(nod))
}

// eg: go test -v  -run=swapList
