package algorithm

// 每一种数字开头的排列有(n-1)!
// 1234
// 1243
// 1324
// 1342
// 1423
// 1432

// 2134
// 2143
// 2314
// 2341
// 2413
// 2431
// n=4, k=10
func getPermutation(n int, k int) string {
	if k == 1 && n == 1 {
		return "1"
	}

	k = k - 1 // 从0开始排

	res := []byte{}
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}

	for len(nums) > 0 {
		once := jiecheng(len(nums) - 1)
		skip := k / once

		res = append(res, '0'+uint8(nums[skip]))
		copy(nums[skip:], nums[skip+1:])
		nums = nums[:len(nums)-1]
		k = k % once
	}

	return string(res)
}
