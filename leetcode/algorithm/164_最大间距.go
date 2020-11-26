package algorithm

import "sort"

// 输入: [3,6,9,1]
// 输出: 3
func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	max := 0
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if max < diff {
			max = diff
		}
	}
	return max
}
