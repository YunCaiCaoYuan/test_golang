package valid_BST

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

func isValidBST(root *TreeNode) bool {
	return validBST(root, math.MinInt64, math.MaxInt64)
}

func validBST(root *TreeNode, min, max int) bool {
	if root != nil {
		fmt.Println("root.Val:", root.Val, "min:", min, "max:", max)
	} else {
		fmt.Println("root:", root, "min:", min, "max:", max)
	}
	if root == nil {
		return true
	}
	if root.Val <= min {
		return false
	}
	if root.Val >= max {
		return false
	}

	return validBST(root.Left, min, root.Val) && validBST(root.Right, root.Val, max)
}

// 输入：root = [2,1,3]
// 输出：true
func Test_isValidBST(t *testing.T) {
	left := &TreeNode{Val: 1}
	right := &TreeNode{Val: 3}
	root := &TreeNode{Val: 2, Left: left, Right: right}

	ret := isValidBST(root)
	fmt.Println(ret)
}

// 输入：root = [5,1,4,null,null,3,6]
// 输出：false
func Test_isValidBST2(t *testing.T) {
	left := &TreeNode{Val: 1}
	rightRoot := &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}
	root := &TreeNode{Val: 5, Left: left, Right: rightRoot}

	ret := isValidBST(root)
	fmt.Println(ret)
}
