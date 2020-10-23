package algorithm

import (
	"sort"
)

// -1,-1,-1,0,0,0,1,1,2
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i >= 1 && nums[i] == nums[i-1] { // 如果后续有相同的就不枚举了
			continue
		}

		start, end := i+1, len(nums)-1
		for end > start {
			if nums[i]+nums[start]+nums[end] == 0 {
				result = append(result, []int{nums[i], nums[start], nums[end]})
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
			} else if nums[i]+nums[start]+nums[end] < 0 {
				start++
			} else if nums[i]+nums[start]+nums[end] > 0 {
				end--
			}
		}
	}
	return result
}
