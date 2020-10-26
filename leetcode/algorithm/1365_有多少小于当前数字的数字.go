package algorithm

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	for i := range nums {
		for j := range nums {
			if i == j {
				continue
			}
			if nums[i] > nums[j] {
				result[i]++
			}
		}
	}
	return result
}

func smallerNumbersThanCurrent1(nums []int) []int {
	l := make([]int, 101)
	for i := range nums {
		l[nums[i]]++
	}

	for i := 1; i < len(l)-1; i++ {
		l[i] = l[i] + l[i-1]
	}

	result := make([]int, len(nums))
	for i := range nums {
		if nums[i] == 0 {
			continue
		}
		result[i] = l[nums[i]-1]
	}

	return result
}
