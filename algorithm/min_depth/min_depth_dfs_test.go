package min_depth

import (
	"fmt"
	"math"
	"testing"
)

// 计算二叉树的最小深度

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 深度优先
func minDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepthDFS(root.Left)
	right := minDepthDFS(root.Right)

	fmt.Println("left:", left, "right:", right, "val:", root.Val)

	if left == 0 || right == 0 {
		return left + right + 1
	}

	return int(math.Min(float64(left), float64(right)) + 1)
}

//  3
// 9 20
//  15 7
func Test_minDepthDFS_1(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	res := minDepthDFS(root)
	fmt.Println(res)
}

// 2
