package algorithm

//输入：obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//输出：2
//解释：
//3x3 网格的正中间有一个障碍物。
//从左上角到右下角一共有 2 条不同的路径：
//1. 向右 -> 向右 -> 向下 -> 向下
//2. 向下 -> 向下 -> 向右 -> 向右

// 走到x,y的格子有几种走法 dp[x][y] = dp[x-1][y] + dp[x][y-1]
// 注意判断边界与上一步是否是障碍物
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 || obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	}
	ans := make([][]int, len(obstacleGrid))
	col, row := false, false
	for i := range ans {
		ans[i] = make([]int, len(obstacleGrid[0]))
		if obstacleGrid[i][0] == 1 {
			col = true
		}
		if !col {
			ans[i][0] = 1
		}
	}
	for i := range ans[0] {
		if obstacleGrid[0][i] == 1 {
			row = true
		}
		if !row {
			ans[0][i] = 1
		}
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i-1][j] == 1 && obstacleGrid[i][j-1] == 1 {
				ans[i][j] = 0
			} else if obstacleGrid[i-1][j] == 1 {
				ans[i][j] = ans[i][j-1]
			} else if obstacleGrid[i][j-1] == 1 {
				ans[i][j] = ans[i-1][j]
			} else {
				ans[i][j] = ans[i-1][j] + ans[i][j-1]
			}
		}
	}
	return ans[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
