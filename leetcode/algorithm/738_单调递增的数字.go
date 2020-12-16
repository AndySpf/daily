package algorithm

import (
	"math"
)

// 输入: N = 232
// 输出: 299
func monotoneIncreasingDigits(N int) int {
	if N < 10 {
		return N
	}
	divisor := 10
	nums := []int{} // 入栈
	for {
		tmp := (N % divisor) / (divisor / 10)
		nums = append(nums, tmp)
		if end := N / divisor; end == 0 {
			break
		}
		divisor *= 10
	}

	newNum := 0
	lowerNine := false
	for i := len(nums) - 1; i >= 0; i-- {
		if lowerNine == true {
			nums[i] = 9
		}

		// 判断什么条件下，开始转换(更好的方案是一次遍历，发现大于等于且减一后仍然大于等于则转换)
		if i > 0 && !lowerNine {
			for j := i; j >= 0; j-- {
				if nums[j] == nums[i] { // 一直相等则继续遍历
					continue
				}
				if nums[j] < nums[i] { // 不相等则判断，如果不满足增序，就开始转换
					lowerNine = true
					nums[i]--
				}
				break
			}
		}
		newNum += nums[i] * int(math.Pow10(i))
	}
	return newNum
}
