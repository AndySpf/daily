package algorithm

// 递归有点没理解正确没搞出来
func PredictTheWinner(nums []int) bool {
	return total(nums, 0, len(nums)-1, 1) >= 0
}

func total(nums []int, start, end int, turn int) int {
	if start == end {
		return nums[start] * turn
	}
	scoreStart := nums[start]*turn + total(nums, start+1, end, -turn)
	scoreEnd := nums[end]*turn + total(nums, start, end-1, -turn)
	return maxScore(scoreStart*turn, scoreEnd*turn) * turn
}

func maxScore(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 当前玩家和另一个玩家最大分数的差值，按照题目规定第一个取的是先手。则对于[0:len(nums)-1] 当前玩家为先手玩家，对于[1:len(nums)-1],则当前玩家变为后手玩家
// dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
func PredictTheWinner1(nums []int) bool {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums))
		dp[i][i] = nums[i]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			dp[i][j] = maxScore(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
		}
	}

	return dp[0][len(nums)-1] >= 0
}
