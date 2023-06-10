package algorithm

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	max := 0
	for i := range height[:len(height)-1] {
		tmp := getMax(height, i)
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func getMax(height []int, index int) int {
	start := height[index]
	area := 0
	for i := index + 1; i < len(height); i++ {
		tmp := minNum(start, height[i]) * (i - index)
		fmt.Println(math.Abs(float64(height[i]-start)), "|", i-index, "|", tmp)
		if tmp > area {
			area = tmp
		}
	}
	return area
}

// 双指针分别指向首尾：minFindRotateSteps(height[0],height[len(height)-1]) * (len(height)-1-0)
// 移动min(height[0],height[len(height)-1])中较小的一个。才可能获得更大结果
// 如果移动较大的那一个，则min求得的结果一定不会更大，只有可能更小。
func maxArea1(height []int) int {
	left, right := 0, len(height)-1
	area := 0

	for left < right {

		tmp := minNum(height[left], height[right]) * (right - left)
		if tmp > area {
			area = tmp
		}
		if height[left] <= height[right] {
			left++
		} else {
			height[right]--
		}
	}
	return area
}
