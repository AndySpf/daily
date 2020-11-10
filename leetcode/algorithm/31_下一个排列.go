package algorithm

// 1,3,5,4,2 => nums[i]==5, target=nums[i-1]=3 => nums[j]==4 =>
// reverse => 1,3,2,4,5 => exchange => 1,4,2,3,5
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
			reverseInts(nums[i:])               // i之后的内容本来就是升序的，此时将这升序翻转就变成了最小的后缀
			j = i + len(nums[i:]) - (j - i) - 1 // 翻转后j的位置，然后将i-1，j交换即可
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
