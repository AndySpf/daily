package algorithm

func isValidSudoku(board [][]byte) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}
			if !checkSudoku(i, j, board) {
				return false
			}
		}
	}
	return true
}

// (2,2) -> 1
// (2,3) -> 1
func checkSudoku(row, col int, board [][]byte) bool {
	item := board[row][col]
	block := (row / 3 * 3) + col/3
	for i := range board[row] {
		if board[row][i] == item && i != col {
			return false
		}
	}
	for i := range board {
		if board[i][col] == item && i != row {
			return false
		}
		continue
	}
	for i := range board {
		for j := range board[i] {
			if (i/3*3)+j/3 != block {
				continue
			}
			if board[i][j] == item && i != row && j != col {
				return false
			}
		}
	}
	return true
}
