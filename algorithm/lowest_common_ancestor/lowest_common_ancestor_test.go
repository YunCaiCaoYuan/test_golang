package lowest_common_ancestor

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

//  二叉树的最近公共祖先
// 递归解法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

//输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
//输出：3
func Test_lowestCommonAncestor(t *testing.T) {
	left := &TreeNode{Val: 5, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}}}
	right := &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 8}}
	root := &TreeNode{Val: 3, Left: left, Right: right}
	ret := lowestCommonAncestor(root, left, right)
	fmt.Println(ret)
}

// 二叉搜索树的最近公共祖先
func lowestCommonAncestorV2(root, p, q *TreeNode) *TreeNode {
	if p.Val < root.Val && root.Val > q.Val {
		return lowestCommonAncestorV2(root.Left, p, q)
	}
	if p.Val > root.Val && root.Val < q.Val {
		return lowestCommonAncestorV2(root.Right, p, q)
	}
	return root
}

func Test_lowestCommonAncestorV2(t *testing.T) {
	left := &TreeNode{Val: 1}
	right := &TreeNode{Val: 3}
	root := &TreeNode{Val: 2, Left: left, Right: right}
	ret := lowestCommonAncestorV2(root, left, right)
	fmt.Println(ret)
}

//    6
//  2   8
// 0 4 7 9
//  3 5
//输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
//输出: 6
func Test_lowestCommonAncestorV2_B(t *testing.T) {
	left := &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}
	right := &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}
	root := &TreeNode{Val: 6, Left: left, Right: right}
	ret := lowestCommonAncestorV2(root, left, right)
	fmt.Println(ret)
}
