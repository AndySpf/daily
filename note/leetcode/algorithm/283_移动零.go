package algorithm

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	slow, fast := 0, 1
	for fast < len(nums) {
		if nums[slow] == 0 && nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
			fast++
			continue
		}

		if nums[slow] != 0 {
			slow++
		}
		fast++
	}
}
