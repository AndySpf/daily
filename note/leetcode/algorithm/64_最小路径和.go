package algorithm

// dp[i][j] 移动到当前格子的最小花费
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		if i == 0 {
			ans[i][0] = grid[i][0]
			continue
		}
		ans[i][0] = ans[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		ans[0][j] = ans[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			ans[i][j] = minNum(ans[i-1][j], ans[i][j-1]) + grid[i][j]
		}
	}
	return ans[m-1][n-1]
}
