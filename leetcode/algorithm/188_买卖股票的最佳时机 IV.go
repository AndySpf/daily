package algorithm

// k = 2, prices = [3,2,6,5,0,3]
// 7

// dp[i][k][0]  第i天，进行k次交易后且未持有任何股票的最大利润
// dp[i][k][1]  第i天，进行k次交易后且持有一支股票的最大利润

// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])   今天手里没股票，则前一天要么手里没股票，休息一天；要么手里有股票，卖掉
// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
func maxProfit4(k int, prices []int) int {
	dp := make([][][]int, len(prices))

	for i := range prices {
		dp[i] = make([][]int, k)
		for j := 0; j < k; j++ {
			dp[i][j] = []int{0, 1}
			if i == 0 {
				dp[0][j][0] = 0
				dp[0][j][1] = -prices[0]
			}
		}
	}

	for i := range prices {
		for j := 0; j < k; j++ {
			dp[i][j][0] = maxNum(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][k][1] = maxNum(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][k][0]
}
