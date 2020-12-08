package algorithm

// 示例:
//
//输入: [-2,1,-3,4,-1,2,1,-5,4]
//输出: 6
//解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
func maxSubArray(nums []int) int {
	max := -(1 << 31)
	count := 0
	for i := range nums {
		if count+nums[i] < 0 {
			if nums[i] > max {
				max = nums[i]
			}
			count = 0
			continue
		}

		count += nums[i]
		if count > max {
			max = count
		}
	}
	return max
}

// 官方分治解法：
func maxSubArray1(nums []int) int {
	return get(nums, 0, len(nums)-1).mSum
}

func pushUp(l, r Status) Status {
	iSum := l.iSum + r.iSum
	lSum := maxMaxSubArray(l.lSum, l.iSum+r.lSum)
	rSum := maxMaxSubArray(r.rSum, r.iSum+l.rSum)
	mSum := maxMaxSubArray(maxMaxSubArray(l.mSum, r.mSum), l.rSum+r.lSum)
	return Status{lSum, rSum, mSum, iSum}
}

func get(nums []int, l, r int) Status {
	if l == r {
		return Status{nums[l], nums[l], nums[l], nums[l]}
	}
	m := (l + r) >> 1
	lSub := get(nums, l, m)
	rSub := get(nums, m+1, r)
	return pushUp(lSub, rSub)
}

func maxMaxSubArray(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Status struct {
	lSum, rSum, mSum, iSum int
}
