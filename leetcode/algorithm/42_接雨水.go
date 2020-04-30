package algorithm

func minNum(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func maxNum(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// [1,3,4,6,5,-1,0]

// [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
// 找到每一个元素可能存储水量的大小，该大小由当前元素左侧最大值和右侧最大值中的较小值减去当前元素高度得到
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	sum := 0
	leftSize, rightSize := make([]int, len(height)), make([]int, len(height))

	leftSize[0] = 0
	leftSize[1] = height[0]
	for i := 2; i < len(height); i++ {
		leftSize[i] = maxNum(leftSize[i-1], height[i-1])
	}

	rightSize[len(height)-1] = 0
	rightSize[len(height)-2] = height[len(height)-1]
	for i := len(height) - 3; i >= 0; i-- {
		rightSize[i] = maxNum(rightSize[i+1], height[i+1])
	}

	for i := range height {
		item := minNum(rightSize[i], leftSize[i]) - height[i]
		if item > 0 {
			sum += item
		}
	}
	return sum
}

// [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
// 左右双指针思路（官方思路）
func trap1(height []int) int {
	if len(height) < 3 {
		return 0
	}
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	sum := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				sum += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				sum += rightMax - height[right]
			}
			right--
		}
	}
	return sum
}
