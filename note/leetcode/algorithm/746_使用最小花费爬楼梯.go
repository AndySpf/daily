package algorithm

// 输入: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
// 输出: 6
// 解释: 最低花费方式是从cost[0]开始，逐个经过那些1，跳过cost[3]，一共花费6。
// dp[i] 前i个楼梯消耗的最少体力
// dp[i] = min(dp[i-1], dp[i-2]]) + cost[i]（登顶那个台阶认为是0即可）
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i <= len(cost); i++ {
		dp[i] = minNum(dp[i-1], dp[i-2])
		if i != len(cost) {
			dp[i] += cost[i]
		}
	}

	return dp[len(cost)]
}
