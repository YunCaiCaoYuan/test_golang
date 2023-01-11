package level_order

import (
	"fmt"
	"testing"
)

type levelOrderDFS struct {
	result [][]int
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (this *levelOrderDFS) levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	this.result = make([][]int, 0)
	this._dfs(root, 0)
	return this.result
}
func (this *levelOrderDFS) _dfs(node *TreeNode, level int) {
	if node == nil {
		return
	}

	if len(this.result) < level+1 {
		this.result = append(this.result, make([]int, 0))
	}

	fmt.Println("level:", level, "val:", node.Val)
	this.result[level] = append(this.result[level], node.Val)

	this._dfs(node.Left, level+1)
	this._dfs(node.Right, level+1)
}

//  3
// 9 20
//  15 7
func Test_levelOrderDFS_1(t *testing.T) {
	lo := new(levelOrderDFS)
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	res := lo.levelOrder(root)
	printArr2(res)
}

//  3
//9  20
// 15 7

func printArr2(list [][]int) {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[i]); j++ {
			fmt.Printf("%d ", list[i][j])
		}
		fmt.Println("")
	}
}
