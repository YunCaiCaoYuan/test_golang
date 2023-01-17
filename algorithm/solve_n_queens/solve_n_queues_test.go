package solve_n_queens

import (
	"fmt"
	"strings"
	"testing"
)

// 在纸上运行一遍

// 回溯
func solveNQueens(n int) [][]string {
	result := make([][]string, 0)

	var backtracking func(int, int, []string)

	backtracking = func(n int, row int, chessboard []string) {
		if row == n {
			fmt.Println("产生,收集结果:")
			printArr(chessboard)
			result = append(result, deepCopyArr(chessboard))
			return
		}

		for col := 0; col < n; col++ {
			//fmt.Println("row:", row, "col:", col)
			if isValid(row, col, chessboard, n) {
				//chessboard[row][col] = 'Q'	// 放置皇后
				chessboard[row] = modifyStr(chessboard[row], col, 'Q')
				//fmt.Println("放置后:", chessboard[row])
				backtracking(n, row+1, chessboard)
				chessboard[row] = modifyStr(chessboard[row], col, '.')
				//fmt.Println("撤销后:", chessboard[row])
				//chessboard[row][col] = '.'	// 回溯，撤销皇后
			}
		}
	}

	chessboard := make([]string, n, n)
	for idx, _ := range chessboard {
		chessboard[idx] = strings.Repeat(".", n)
	}
	backtracking(n, 0, chessboard)
	return result
}

func isValid(row int, col int, chessboard []string, n int) bool {
	// 检查列
	for i := 0; i < row; i++ { // 这是一个剪枝 ???
		if chessboard[i][col] == 'Q' {
			return false
		}
	}
	// 检查45度角是否有皇后
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if chessboard[i][j] == 'Q' {
			return false
		}
	}
	// 检查135度角是否有皇后
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if chessboard[i][j] == 'Q' {
			return false
		}
	}
	return true
}

func modifyStr(str string, i int, modify byte) string {
	s := []byte(str)
	s[i] = modify
	return string(s)
}

func deepCopyArr(oldList []string) []string {
	list := make([]string, 0, len(oldList))
	for _, item := range oldList {
		list = append(list, item)
	}
	return list
}

func Test_SolveNQueues(t *testing.T) {
	ret := solveNQueens(4)
	printStrArr(ret)
}

// [".Q..","...Q","Q...","..Q."]
// ["..Q.","Q...","...Q",".Q.."]

func printStrArr(list [][]string) {
	//for i := 0; i < len(list); i++ {
	//	for j := 0; j < len(list[i]); j++ {
	//		fmt.Printf("%s ", list[i][j])
	//	}
	//	fmt.Println("")
	//}
}
func printArr(list []string) {
	for j := 0; j < len(list); j++ {
		fmt.Printf("%s ", list[j])
	}
	fmt.Println("")
}

func Test_SolveNQueues_2(t *testing.T) {
	ret := solveNQueens(1)
	printStrArr(ret)
}

// Q

func Test_SolveNQueues_3(t *testing.T) {
	ret := solveNQueens(2)
	printStrArr(ret)
}
func Test_SolveNQueues_4(t *testing.T) {
	ret := solveNQueens(3)
	printStrArr(ret)
}
