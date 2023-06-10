package algorithm

func searchInsert(nums []int, target int) int {
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	if target < nums[0] {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}

		if i < len(nums)-1 {
			if nums[i] < target && nums[i+1] > target {
				return i + 1
			}
		}
	}
	return -1
}

// 因为是有序数组，可以采用二分法
func searchInsert1(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1
	ans := n
	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
