package algorithm

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	for i := 1; i < len(nums); {
		if nums[i] == nums[i-1] {
			copy(nums[i-1:], nums[i:])
			nums = nums[:len(nums)-1]
		} else {
			i++
		}
	}
	return len(nums)
}

// 快慢指针,用快指针跳过重复项，慢指针给各个位上赋值不重复的数据  1,1,3,4
func removeDuplicates1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	fast, slow := 1, 0
	for fast < len(nums) {
		if nums[fast] == nums[slow] {
			fast++
		} else {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
