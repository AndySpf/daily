package dynamic_programming

// coins 中为硬币面值， amount为要凑的总量。求凑出amount所需要的最少硬币个数
// coins: [1, 2, 5]; amount: 18
// dp[18] = min(dp[13], dp[16], dp[17]) + 1
// 注意边界判定

func chipInCoins(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	dp := make([]int, amount+1)

	for i := range coins {
		if len(dp) > coins[i] {
			dp[coins[i]] = 1
		}
	}

	for i := 1; i <= amount; i++ {
		min := 0
		for j := range coins {
			if i-coins[j] < 0 || dp[i-coins[j]] == 0 { // 某一数小于硬币面值或者某一个子问题无解则跳过
				continue
			}

			if min > dp[i-coins[j]] || min == 0 {
				min = dp[i-coins[j]]
			}
		}

		if min != 0 {
			dp[i] = min + 1
		}
	}

	if dp[amount] == 0 {
		return -1
	}

	return dp[amount]
}
