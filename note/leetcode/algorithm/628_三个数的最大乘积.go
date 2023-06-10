package algorithm

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)

	l := len(nums)
	res := nums[l-1] * nums[l-2] * nums[l-3]
	if nums[0] < 0 && nums[1] < 0 {
		tmp1 := -nums[0] * -nums[1] * nums[len(nums)-1]
		if tmp1 > res {
			res = tmp1
		}
	}
	return res
}
