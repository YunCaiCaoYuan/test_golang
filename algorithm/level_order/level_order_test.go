package level_order

import (
	"fmt"
	"testing"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bfs
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	result := make([][]int, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		level_size := len(queue)
		current_level := make([]int, 0)

		for i := 0; i < level_size; i++ {
			node := queue[0]
			queue = queue[1:]
			current_level = append(current_level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		printArr(current_level)
		result = append(result, current_level)
	}

	return result
}

func printArr(list []int) {
	for j := 0; j < len(list); j++ {
		fmt.Printf("%d ", list[j])
	}
	fmt.Println("")
}

func printArr2(list [][]int) {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			fmt.Printf("%d ", list[i][j])
		}
		fmt.Println("")
	}
}

func Test_levelOrder_1(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	res := levelOrder(root)
	printArr2(res)
}
