package algorithm

//给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。
//
//返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
func findTargetSumWays(nums []int, S int) int {
	var count int
	backtrack(&count, nums, 0, 0, S)
	return count
}

func backtrack(count *int, nums []int, i int, res int, S int) {
	if i == len(nums) {
		if res == S {
			*count++
		}
		return
	}
	backtrack(count, nums, i+1, res+nums[i], S)
	backtrack(count, nums, i+1, res-nums[i], S)
}

// dp[i][j+1000] = dp[i-1][j+1000-nums[i]] + dp[i-1][j+1000+nums[i]] ->

// dp[i][j+1000] = dp[i][j+1000] + dp[i-1][j+1000-nums[i]]
// dp[i][j+1000] = dp[i][j+1000] + dp[i-1][j+1000+nums[i]]
func findTargetSumWays1(nums []int, S int) int {
	if len(nums) == 0 {
		if S == 0 {
			return 1
		} else {
			return 0
		}
	}

	if S > 1000 {
		return 0
	}

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, 2000)
	}

	dp[0][nums[0]+1000] = 1   //nums中只有一项,且目标为这一项的正数
	dp[0][-nums[0]+1000] += 1 //nums中只有一项,且目标为这一项的负数。且nums[0]可能为0,所以要nums[i]==-nums[i]，即凑出0有两种方案,因此要用+=符号

	for i := 1; i < len(nums); i++ {
		for j := -1000; j < 1000; j++ {
			j1 := j - nums[i]
			j2 := j + nums[i]
			if j1 < 1000 && j1 > -1000 {
				dp[i][j+1000] = dp[i][j+1000] + dp[i-1][j1+1000]
			}
			if j2 < 1000 && j2 > -1000 {
				dp[i][j+1000] = dp[i][j+1000] + dp[i-1][j2+1000]
			}
		}
	}
	return dp[len(nums)-1][S+1000]
}

// sum(A) - sum(B) = target
// 2sum(A) = target + sum(nums)
// 存在哪些子集，使得子集内的和为 sum(A) = (sum(nums) + S) / 2
func findTargetSumWays2(nums []int, S int) int {
	return 0
}
