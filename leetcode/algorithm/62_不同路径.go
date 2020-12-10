package algorithm

import (
	"math/big"
)

//示例 1:
//输入: m = 3, n = 2
//输出: 3
//解释:
//从左上角开始，总共有 3 条路径可以到达右下角。
//1. 向右 -> 向右 -> 向下
//2. 向右 -> 向下 -> 向右
//3. 向下 -> 向右 -> 向右
//示例 2:
//输入: m = 7, n = 3
//输出: 28

// 排列组合
func uniquePaths(m int, n int) int {
	// 总移动步数中取出向下移动次数的组合方案（有序）
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}

// dp[i][j] 有多少种走法
// dp[i][j] = dp[i-1][j] + dp[i][j-1]
func uniquePaths1(m int, n int) int {
	dp := make([][]int, m)

	// init
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := range dp[0] {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
