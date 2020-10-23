package algorithm

/*
一个数独的解法需遵循如下规则：
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
*/
/*
	5  3  .  .  7  .  .  .  .
	6  .  .  1  9  5  .  .  .
	.  9  8  .  .  .  .  6  .
	8  .  .  .  6  .  .  .  3
	4  .  .  8  .  3  .  .  1
	7  .  .  .  2  .  .  .  6
	.  6  .  .  .  .  2  8  .
	.  .  .  4  1  9  .  .  5
	.  .  .  .  8  .  .  7  9
*/

// 回溯算法 https://github.com/labuladong/fucking-algorithm/blob/master/%E7%AE%97%E6%B3%95%E6%80%9D%E7%BB%B4%E7%B3%BB%E5%88%97/%E5%9B%9E%E6%BA%AF%E7%AE%97%E6%B3%95%E8%AF%A6%E8%A7%A3%E4%BF%AE%E8%AE%A2%E7%89%88.md
func solveSudoku(board [][]byte) {
	row := [9][9]bool{}
	col := [9][9]bool{}
	box := [9][9]bool{}
	for rowK := 0; rowK < 9; rowK++ {
		for colK := 0; colK < 9; colK++ {
			if board[rowK][colK] != '.' {
				num := board[rowK][colK] - '1'
				boxK := (rowK/3)*3 + colK/3
				row[rowK][num], col[colK][num], box[boxK][num] = true, true, true
			}
		}
	}
	fill(board, 0, row, col, box)
}

func fill(board [][]byte, n int, row [9][9]bool, col [9][9]bool, box [9][9]bool) bool {
	if n == 81 {
		return true
	}
	rowK := n / 9
	colK := n % 9
	if board[rowK][colK] != '.' {
		return fill(board, n+1, row, col, box)
	}

	boxK := (rowK/3)*3 + colK/3
	for num := 0; num < 9; num++ {
		if !row[rowK][num] && !col[colK][num] && !box[boxK][num] {
			board[rowK][colK] = byte('1' + num)
			row[rowK][num], col[colK][num], box[boxK][num] = true, true, true
			if fill(board, n+1, row, col, box) { //下一个填充
				return true
			}
			row[rowK][num], col[colK][num], box[boxK][num] = false, false, false //失败回溯
		}
	}
	board[rowK][colK] = '.'
	return false
}

// 第二次写
func solveSudoku1(board [][]byte) {
	rows := [9][9]bool{} // 第n行是否已经存在某个数字
	cols := [9][9]bool{}
	boxs := [9][9]bool{}
	for rowIndex := 0; rowIndex < 9; rowIndex++ {
		for colIndex := 0; colIndex < 9; colIndex++ {
			if board[rowIndex][colIndex] != '.' {
				num := board[rowIndex][colIndex] - '1' // 减'0'后是1-9，因此减'1'
				rows[rowIndex][num], cols[colIndex][num] = true, true
				boxIndex := (rowIndex/3)*3 + colIndex/3
				boxs[boxIndex][num] = true
			}
		}
	}
	fillNum(board, 0, rows, cols, boxs)
}

func fillNum(board [][]byte, n int, rows [9][9]bool, cols [9][9]bool, boxs [9][9]bool) bool {
	if n == 81 {
		return true
	}

	rowIndex := n / 9
	colIndex := n % 9

	if board[rowIndex][colIndex] != '.' {
		return fillNum(board, n+1, rows, cols, boxs)
	}
	boxIndex := (rowIndex/3)*3 + colIndex/3

	for number := 0; number < 9; number++ {
		if !rows[rowIndex][number] && !cols[colIndex][number] && !boxs[boxIndex][number] {
			board[rowIndex][colIndex] = uint8(number) + '1'
			rows[rowIndex][number] = true
			cols[colIndex][number] = true
			boxs[boxIndex][number] = true

			if fillNum(board, n+1, rows, cols, boxs) {
				return true
			}
			rows[rowIndex][number] = false
			cols[colIndex][number] = false
			boxs[boxIndex][number] = false
			board[rowIndex][colIndex] = '.'
		}
	}

	return false
}

// 第三次写
func solveSudoku3(board [][]byte) {
	status := map[string][][]bool{}
	status["row"] = make([][]bool, 9)
	status["col"] = make([][]bool, 9)
	status["block"] = make([][]bool, 9)

	for i := 0; i < 9; i++ {
		status["row"][i] = make([]bool, 9)
		status["col"][i] = make([]bool, 9)
		status["block"][i] = make([]bool, 9)
	}

	for i := range board {
		for j := range board[i] {
			if num := board[i][j]; num != '.' {
				block := i/3*3 + j/3
				status["row"][i][num-'0'-1] = true
				status["col"][j][num-'0'-1] = true
				status["block"][block][num-'0'-1] = true
			}
		}
	}

	fillNum3(0, board, status)
}

func fillNum3(n int, board [][]byte, status map[string][][]bool) bool {
	if n == 9*9 {
		return true
	}

	col := n % 9
	row := n / 9
	block := row/3*3 + col/3
	if num := board[row][col]; num != '.' {

		return fillNum3(n+1, board, status)
	}

	for num := 1; num <= 9; num++ {
		if !status["row"][row][num-1] && !status["col"][col][num-1] && !status["block"][block][num-1] {

			status["row"][row][num-1] = true
			status["col"][col][num-1] = true
			status["block"][block][num-1] = true
			board[row][col] = uint8(num) + '0'
			if fillNum3(n+1, board, status) == true {
				return true
			}
			board[row][col] = '.'
			status["row"][row][num-1] = false
			status["col"][col][num-1] = false
			status["block"][block][num-1] = false
		}
	}
	return false
}
