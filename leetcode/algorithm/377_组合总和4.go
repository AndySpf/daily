package algorithm

import "fmt"

/*
给定一个由正整数组成且不存在重复数字的数组，找出和为给定目标正整数的组合的个数。

示例:

nums = [1, 2, 3]
target = 4
所有可能的组合为：
(1, 1, 1, 1)
(1, 1, 2)
(1, 2, 1)
(1, 3)
(2, 1, 1)
(2, 2)
(3, 1)
*/

var totalCombinationSum4 int

// 数学法，int会超限
func combinationSum4(candidates []int, target int) int {
	targetCombinationSum = target
	allCombinationSum = candidates

	totalCombinationSum4 = 0

	backCombinationSum4(0, []int{})
	return totalCombinationSum4
}

func backCombinationSum4(index int, item []int) {
	v := sumInt(item)
	if v >= targetCombinationSum {
		if v == targetCombinationSum {
			fmt.Println(item)
			totalCombinationSum4 += getRes(item)
		}
		return
	}

	for i := index; i < len(allCombinationSum); i++ {
		if allCombinationSum[i] > targetCombinationSum {
			continue
		}
		item = append(item, allCombinationSum[i])
		backCombinationSum4(i, item)
		item = item[:len(item)-1]
	}
}

func jiecheng(num int) int {
	res := 1
	for i := 2; i <= num; i++ {
		res = res * i
		fmt.Println(res)
	}
	return res
}

func getRes(item []int) int {
	m := map[int]int{}
	for i := range item {
		if _, ok := m[item[i]]; ok {
			m[item[i]]++
			continue
		}
		m[item[i]] = 1
	}
	all := jiecheng(len(item))
	tmp := 1
	for i := range m {
		tmp = tmp * jiecheng(m[i])
	}
	return all / tmp
}

// 凑齐i有多少中方法
// dp[i] = dp[i-nums[j]] + dp[i]   // 可以选择第j个元素，也可以不选择第j个元素
func combinationSum4_1(candidates []int, target int) int {
	var dp = map[int]int{}
	dp[0] = 1
	for t := 1; t <= target; t++ {
		for index := range candidates {
			if t < candidates[index] {
				continue
			}
			dp[t] = dp[t] + dp[t-candidates[index]]
		}
	}
	return dp[target]
}
