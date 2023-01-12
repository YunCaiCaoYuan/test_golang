package generate_parenthesis

import (
	"fmt"
	"testing"
)

// n代表生成括号的对数
// 设计用于能够生成所有可能的并且有效的括号组合

// 递归+剪枝
func generateParenthesis(n int) []string {

	list := make([]string, 0)

	var _gen func(int, int, int, string)
	_gen = func(left int, right int, n int, result string) {
		if left == n && right == n {
			list = append(list, result)
			return
		}
		if left < n {
			_gen(left+1, right, n, result+"(")
		}
		if left > right && right < n {
			_gen(left, right+1, n, result+")")
		}
	}

	_gen(0, 0, n, "")

	return list
}

func Test_generateParenthesis(t *testing.T) {
	res := generateParenthesis(3)
	printArr(res)
}
func Test_generateParenthesis_2(t *testing.T) {
	res := generateParenthesis(4)
	printArr(res)
}

func printArr(list []string) {
	for _, item := range list {
		fmt.Println(item)
	}
}
