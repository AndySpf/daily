package algorithm

// 1,2,3 => 1,3,2
// 3,2,1 => 1,2,3
// 1,3,4,2
// 1,3,5,4,2  2,3,1
func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			target := nums[i-1]

			j := i // 找到刚好比target大的位置
			for j < len(nums)-1 {
				if nums[j+1] > target {
					j++
				} else {
					break
				}
			}
			reverseInts(nums[i:])
			j = i + len(nums[i:]) - (j - i) - 1 // 翻转后j的位置
			nums[j], nums[i-1] = nums[i-1], nums[j]
			return
		}
	}
	reverseInts(nums)
	return
}

func reverseInts(sli []int) {
	for i := len(sli)/2 - 1; i >= 0; i-- {
		j := len(sli) - i - 1
		sli[j], sli[i] = sli[i], sli[j]
	}
}
