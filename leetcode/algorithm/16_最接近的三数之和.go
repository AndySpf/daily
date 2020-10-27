package algorithm

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	minDiff := int(^uint32(0))
	result := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i >= 1 && nums[i] == nums[i-1] { // 如果后续有相同的就不枚举了
			continue
		}

		start, end := i+1, len(nums)-1
		for end > start {
			sum := nums[i] + nums[start] + nums[end]
			diff := int(math.Abs(float64(target - sum)))
			if sum == target {
				return target
			} else if sum < target {
				if minDiff > diff {
					result = sum
					minDiff = diff
				}
				start++
			} else if sum > target {
				if minDiff > diff {
					result = sum
					minDiff = diff
				}
				end--
			}
		}
	}
	return result
}
