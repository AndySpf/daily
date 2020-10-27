package algorithm

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] { // 如果后续有相同的就不枚举了
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] { // 如果后续有相同的就不枚举了
				continue
			}
			start, end := j+1, len(nums)-1
			for end > start {
				if nums[i]+nums[j]+nums[start]+nums[end] == target {
					result = append(result, []int{nums[i], nums[j], nums[start], nums[end]})
					// 如果凑齐0了 则start往前找直到找到一个不相同的数，end往后找，直到找到一个不相同的数。
					// 因为有序且不重复，必须两个同时动
					start++
					for nums[start] == nums[start-1] && start < end {
						start++
					}
					end--
					for nums[end] == nums[end+1] && end > start {
						end--
					}
				} else if nums[i]+nums[j]+nums[start]+nums[end] < target {
					start++
				} else if nums[i]+nums[j]+nums[start]+nums[end] > target {
					end--
				}
			}
		}
	}
	return result
}
