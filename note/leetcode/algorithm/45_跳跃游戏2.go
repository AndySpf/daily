package algorithm

//[3,4,1,3,1,1,7]
var minStep uint32

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minStep = ^uint32(0)
	jumpDfs(nums, 0, 0)
	return int(minStep)
}

func jumpDfs(nums []int, index int, step uint32) {
	if index == len(nums)-1 {
		if step < minStep {
			minStep = step
		}
		return
	}

	if index+nums[index] >= len(nums)-1 {
		if step+1 < minStep {
			minStep = step + 1
		}
		return
	}

	for i := 1; i <= nums[index]; i++ {
		if index+i <= len(nums)-1 {
			if nums[index] > i+nums[i+index] {
				if i+index < len(nums)-1 {
					continue
				}
			}
			jumpDfs(nums, index+i, step+1)
		}
	}
}

// 贪心算法
//[3,4,1,3,1,1,7]
func jump1(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}

	curIndex := 0
	step := 0
	for {
		curIndex = getNextIndex(nums, curIndex)
		if curIndex == -1 {
			return 0
		}
		step++
		if curIndex == len(nums)-1 {
			return step
		}
	}
}

func getNextIndex(nums []int, index int) int {
	v := nums[index]
	maxDistance := 0
	nextIndex := -1
	for i := index + 1; i <= index+v; i++ {
		if i >= len(nums)-1 {
			return len(nums) - 1
		}

		if nums[i]+i >= maxDistance {
			maxDistance = nums[i] + i
			nextIndex = i
		}
	}
	return nextIndex
}
