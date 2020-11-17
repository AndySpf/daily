package algorithm

// 无重复数字的全排列
func permute(nums []int) [][]int {
	permuteRes := make([][]int, 0, jiecheng(len(nums)))

	var dfsPermute func(queue []int)
	dfsPermute = func(queue []int) {
		if len(queue) == len(nums) {
			tmp := make([]int, len(nums))
			copy(tmp, queue)
			permuteRes = append(permuteRes, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if hasInt(queue, nums[i]) {
				continue
			}
			queue = append(queue, nums[i])
			dfsPermute(queue)
			queue = queue[:len(queue)-1]
		}
	}
	queue := make([]int, 0, len(nums))
	dfsPermute(queue)

	return permuteRes
}

func hasInt(sli []int, num int) bool {
	for i := range sli {
		if sli[i] == num {
			return true
		}
	}
	return false
}
