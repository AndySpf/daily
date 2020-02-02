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

var (
	rows   = make([]map[uint8]bool, 9)
	cols   = make([]map[uint8]bool, 9)
	blocks = make([]map[uint8]bool, 9)
)

func solveSudoku(board [][]byte) {
	// 第一行能填写的数字有1,2,4,6,8,9
	// 第三列能填的数字有1,2,3,4,5,6,8,9
	// 第一块能填入的数字有1，2，4，7
	// board（0,2)属于第一行第三列第一块，因此可以选择填入的数字有t[0][2] = [1,2,4]
	// 如果填入了1，则第一行，第三列，第一块的可选数字都要去掉1
	// 然后继续往下走。如果发现没有可选值填入，则回退到上一个数字如(0,2)，
	// 然后尝试填入t[0][2]的下一个数字2，且第一行第三列第一块的可选数字要加上1，去掉2
	// 直到board(8,8)被填入数字

	for i := 0; i < 9; i++ {
		rows[i] = map[uint8]bool{}
		cols[i] = map[uint8]bool{}
		blocks[i] = map[uint8]bool{}
		for index := range board[i] {
			if board[i][index] != '.' {
				rows[i][board[i][index]] = false
			}
		}

		for index := 0; index < 9; index++ {
			if board[index][i] != '.' {
				cols[i][board[index][i]] = false
			}
		}

		// 0->(0,2)(0,2)
		// 1->(3,5)(0,2)
		// 2->(6,8)(0,2)
		// 3->(0,2)(3,5)
		// 4->(3,5)(3,5)
		y := 2*(i%3) + (i % 3)
		x := (i / 3) * 3
		for row := x; row < x+3; row++ {
			for col := y; col < y+3; col++ {
				if board[row][col] != '.' {
					blocks[i][board[row][col]] = false
				}
			}
		}
	}

	fillNum(0, 0, board)

}

func fillNum(i, j int, board [][]byte) bool {
	if i == 8 && j == 9 {
		return true
	}

	if j == 9 {
		j = 0
		i++
	}

	if board[i][j] != '.' {
		j++
		return fillNum(i, j, board)
	}
	isFilled := false
	for num := 1; num < 10; num++ {
		if checkNum(num, i, j) {
			// 检查通过可以使用这个数则对应行和列的要增加一个不可用元素
			rows[i][uint8(num)+uint8('0')] = false
		}
		isFilled = true
		if !fillNum(i, j, board) {
			isFilled = false
			continue
		} else {
			return true
		}
	}
	if !isFilled {
		return false // 返回上一个递归中
	}
	return true
}

func checkNum(num, i, j int) bool {
	if _, ok := rows[i][uint8(num)+uint8('0')]; ok {
		return false
	}
	if _, ok := cols[j][uint8(num)+uint8('0')]; ok {
		return false
	}
	if _, ok := blocks[(j/3)*3+(i/3)][uint8(num)+uint8('0')]; ok {
		return false
	}
	return true
}
