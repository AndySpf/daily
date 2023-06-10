package algorithm

func pivotIndex(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	index := 0
	left, right := 0, sumInt(nums[index+1:])
	for index < len(nums)-1 {
		if left != right {
			left += nums[index]
			right -= nums[index+1]
			index++
		}
		if left == right {
			return index
		}
	}
	return -1
}
