package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func Shownode(p *Node) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.next //移动指针
	}
}

// 尾插法
func main() {
	var head = new(Node)
	head.data = 0
	var tail *Node
	tail = head //tail用于记录最末尾的节点的地址，刚开始tail的的指针指向头节点
	for i := 1; i < 10; i++ {
		var node = Node{data: i}
		(*tail).next = &node
		tail = &node
	}

	Shownode(head) //遍历结果
}
