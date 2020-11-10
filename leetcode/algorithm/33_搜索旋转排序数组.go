package algorithm

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	// 找到旋转点
	start, end := 1, len(nums)-1
	rotate := 0
	for start <= end {
		if nums[start] < nums[start-1] {
			rotate = start
			break
		}
		if nums[end] < nums[end-1] {
			rotate = end
			break
		}
		start++
		end--
	}

	oriLen := len(nums)
	diff := len(nums) - rotate
	// 还原有序数组
	nums = append(nums, nums[:rotate]...)
	nums = nums[rotate:]

	// 二分法查找
	left, right := 0, len(nums)-1
	for left <= right {
		pos := (right-left)/2 + left
		if nums[pos] > target {
			right = pos - 1
		}
		if nums[pos] < target {
			left = pos + 1
		}
		// 3, 1 => 3
		// 3, 4, 1, 2 => 3,4,1,2,3,4
		if nums[pos] == target {
			if pos+rotate < oriLen {
				return pos + rotate
			} else {
				return pos - diff
			}
		}
	}
	return -1
}
