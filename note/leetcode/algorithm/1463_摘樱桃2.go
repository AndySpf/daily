package algorithm

func cherryPickup(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid) == 1 {
		return grid[0][0] + grid[0][len(grid[0])-1]
	}

	getValue := func(row, i, j int) int {
		if i != j {
			return grid[row][i] + grid[row][j]
		}
		return grid[row][i]
	}

	// 记忆表，减少dfs次数
	ans := make([][][]int, len(grid))
	for i := range ans {
		ans[i] = make([][]int, len(grid[0]))
		for j := range ans[i] {
			ans[i][j] = make([]int, len(grid[0]))
			for k := range ans[i][j] {
				ans[i][j][k] = -1
			}
		}
	}

	var dfsCherry func(row, col1, col2 int) int
	dfsCherry = func(row, col1, col2 int) int {
		if row == len(grid)-1 {
			return getValue(row, col1, col2)
		}

		totalCherry := 0
		for i := col1 - 1; i <= col1+1; i++ {
			for j := col2 - 1; j <= col2+1; j++ {
				if i >= 0 && i <= len(grid[0])-1 && j >= 0 && j <= len(grid[0])-1 {
					var next int
					if ans[row+1][i][j] != -1 {
						next = ans[row+1][i][j]
					} else {
						next = dfsCherry(row+1, i, j)
					}
					ans[row+1][i][j] = next
					totalCherry = maxNum(totalCherry, next)
				}
			}
		}
		return totalCherry + getValue(row, col1, col2)
	}
	return dfsCherry(0, 0, len(grid[0])-1)
}
