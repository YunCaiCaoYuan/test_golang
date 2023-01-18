package solve_sudoku

import (
	"fmt"
	"testing"
)

// 解数独
func solveSudoku(board [][]byte) {
	ss := &SolveSudoku{board: board}
	ss.backtracking()
	printArr2(ss.board)
}

type SolveSudoku struct {
	board [][]byte
}

func (this *SolveSudoku) backtracking() bool {
	for i := 0; i < len(this.board); i++ {
		for j := 0; j < len(this.board[0]); j++ {
			if this.board[i][j] != '.' {
				continue
			}
			for k := '1'; k <= '9'; k++ {
				if this.isValid(i, j, byte(k)) {
					this.board[i][j] = byte(k)
					if this.backtracking() {
						return true
					}
					this.board[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
}

func (this *SolveSudoku) isValid(row, col int, val byte) bool {
	for i := 0; i < 9; i++ {
		if this.board[row][i] == val {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if this.board[j][col] == val {
			return false
		}
	}
	startRow := row / 3 * 3
	startCol := col / 3 * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if this.board[i][j] == val {
				return false
			}
		}
	}
	return true
}

func Test_solveSudoku(t *testing.T) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
}

func printArr2(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			fmt.Printf("%s ", string(board[i][j]))
		}
		fmt.Println("")
	}
}
