package max_depth

import (
	"fmt"
	"math"
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

// 深度优先，使用递归来做
func maxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return int(math.Max(float64(maxDepthDFS(root.Left)), float64(maxDepthDFS(root.Right))) + 1)
}

//  3
// 9 20
//  15 7
func Test_maxDepthDFS_1(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	res := maxDepthDFS(root)
	fmt.Println(res)
}

// 3
