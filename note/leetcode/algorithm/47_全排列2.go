package algorithm

import "sort"

// 有重复数字的排列数
// 1. 得出结果后去重
// 2. 数组排序，然后按照官方 if (i > 0 && nums[i] == nums[i - 1] && !vis[i - 1])
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	permuteRes := [][]int{}
	vis := make([]bool, len(nums))
	var dfsPermute func(queue []int)
	dfsPermute = func(queue []int) {
		if len(queue) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, queue)
			permuteRes = append(permuteRes, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			// 只有索引为i的是true，索引为i-1的是false才可以正常添加
			// 这样可以保证如果是相同字符，添加的永远是第一个没用使用的字符
			if vis[i] || (i > 0 && nums[i] == nums[i-1] && !vis[i-1]) {
				continue
			}

			vis[i] = true
			queue = append(queue, nums[i])
			dfsPermute(queue)
			vis[i] = false
			queue = queue[:len(queue)-1]
		}
	}
	queue := make([]int, 0, len(nums))
	dfsPermute(queue)

	return permuteRes
}
