package count_paths

import (
	"fmt"
	"testing"
)

// 时间复杂度：O(  n)
//		       2
func countPaths(grid [][]bool, row, col int) int {
	if !validSquare(grid, row, col) {
		return 0
	}
	if isAtEnd(grid, row, col) {
		return 1
	}
	return countPaths(grid, row+1, col) + countPaths(grid, row, col+1)
}

func validSquare(grid [][]bool, row, col int) bool {
	if row >= len(grid) || col >= len(grid[0]) {
		return false
	}
	return grid[row][col] == false
}

func isAtEnd(grid [][]bool, row, col int) bool {
	return row == len(grid)-1 && col == len(grid[0])-1
}

// 27
func Test_countPaths(t *testing.T) {
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
	ret := countPaths(grid, 0, 0)
	fmt.Println(ret)
}
