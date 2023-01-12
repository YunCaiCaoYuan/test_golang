package max_depth

import (
	"fmt"
	"testing"
)

// bfs
// 地毯式扫荡，第一个叶子节点的深度即为最小深度，扫荡结束即可得到最大深度

// dfs
// max，min，如果是叶子节点则更新这两个值

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先，使用队列来做
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var ans int

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}

//  3
// 9 20
//  15 7
func Test_maxDepth_1(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	res := maxDepth(root)
	fmt.Println(res)
}

// 3
