package count_paths

import (
	"fmt"
	"testing"
)

// 动态规划
// 时间复杂度：O(m*n)
// 空间复杂度：O(m*n)
func countPathsV2(grid [][]bool, row, col int) int {
	opt := make([][]int, len(grid))
	for idx, _ := range grid {
		opt[idx] = make([]int, len(grid))
	}
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			if i == len(grid)-1 || j == len(grid)-1 {
				opt[i][j] = 1
			} else {
				if grid[i][j] == false {
					opt[i][j] = opt[i+1][j] + opt[i][j+1]
				} else {
					opt[i][j] = 0
				}
			}
		}
	}
	//printArr(opt)
	return opt[row][col]
}

func printArr(list [][]int) {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list[0]); j++ {
			fmt.Printf("%2d ", list[i][j])
		}
		fmt.Println("")
	}
}

func Test_countPathsV2(t *testing.T) {
	grid := [][]bool{
		{false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, true, false},
		{false, false, false, false, true, false, false, false},
		{true, false, true, false, false, true, false, false},
		{false, false, true, false, false, false, false, false},
		{false, false, false, true, true, false, true, false},
		{false, true, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, false},
	}
	ret := countPathsV2(grid, 0, 0)
	fmt.Println(ret)
}
