package algorithm

// 输入: nums = [5,7,7,8,8,10], target = 8
// 输出: [3,4]
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	start, end := -1, -1
	for left <= right {
		pos := (left + right) / 2
		if nums[pos] > target {
			right = pos - 1
		}
		if nums[pos] < target {
			left = pos + 1
		}
		if nums[pos] == target {
			start, end = pos, pos
			for {
				if start > 0 && nums[start-1] == target {
					start--
					continue
				}
				if end < len(nums)-1 && nums[end+1] == target {
					end++
					continue
				}
				break
			}
			break
		}
	}
	return []int{start, end}
}
