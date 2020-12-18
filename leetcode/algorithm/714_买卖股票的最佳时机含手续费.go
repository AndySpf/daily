package algorithm

// 示例 1:
//
//输入: prices = [1, 3, 2, 8, 4, 9], fee = 2
//输出: 8
//解释: 能够达到的最大利润:
//在此处买入 prices[0] = 1
//在此处卖出 prices[3] = 8
//在此处买入 prices[4] = 4
//在此处卖出 prices[5] = 9
//总利润: ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
// 4,5,2,4,3,3,1,2,5,4
func maxProfit(prices []int, fee int) int { // ??? 这个思路为什么过不了很长的用例
	prices = append(prices, 0)
	buy, profit := prices[0], 0
	last := []int{-1, -1}
	for i := range prices[:len(prices)-1] {
		if prices[i] < buy {
			buy = prices[i]
		}
		if prices[i] > prices[i+1] && prices[i]-buy > fee {
			if last[0] != -1 {
				if last[1]-buy > fee {
					profit += last[1] - last[0] - fee
					last[0], last[1] = buy, prices[i]
					buy = prices[i+1]
				} else {
					last[1] = prices[i]
				}
			} else {
				last[0], last[1] = buy, prices[i]
				buy = prices[i+1]
			}
		}
	}
	if last[1]-last[0] > fee {
		profit += last[1] - last[0] - fee
	}
	return profit
}

// 4,5,2,4,3,3,1,2,5,4
func maxProfit1(prices []int, fee int) int {
	n := len(prices)
	buy := prices[0] + fee
	profit := 0
	for i := 1; i < n; i++ {
		if prices[i]+fee < buy {
			buy = prices[i] + fee
		} else if prices[i] > buy {
			profit += prices[i] - buy
			buy = prices[i]
		}
	}
	return profit
}
