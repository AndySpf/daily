package algorithm

//输入: nums = [1,3], n = 6
//输出: 1
//解释:
//根据 nums 里现有的组合 [1], [3], [1,3]，可以得出 1, 3, 4。
//现在如果我们将 2 添加到 nums 中， 组合变为: [1], [2], [3], [1,3], [2,3], [1,2,3]。
//其和可以表示数字 1, 2, 3, 4, 5, 6，能够覆盖 [1, 6] 区间里所有的数。
//所以我们最少需要添加一个数字。

// 对于正整数 x，如果区间 [1,x−1] 内的所有数字都已经被覆盖，
// 且 x 在数组中，则区间 [1,2x−1] 内的所有数字也都被覆盖。证明如下
// 1<y<=x-1 => 1<x+y<=2x-1
// [1,3,5,10],20 => 2
// [1,3]

// [1,2)  [1,5)
func minPatches(nums []int, n int) (patches int) {
	for i, x := 0, 1; x <= n; {
		if i < len(nums) && nums[i] < x {
			// 原来可以覆盖的值域为[1,x),因为nums[i]<x，因此nums[i]本身是已经被覆盖了的
			// 因此[0,x+nums[i])一定也是可以覆盖到的，即x+nums[i]成为了新的最小未被覆盖整数
			x += nums[i]
			i++
		} else if i < len(nums) && nums[i] == x {
			// x存在于数组中，此时值域变为[1,x]，根据之前证明可得[1,2x-1)内全部被覆盖，对比数组中下一个
			x *= 2
			i++
		} else { // 如果数组里的值大于了x，则需要额外向数组中添加x，添加后值域变为[1,2x-1)，因数组中当前数字尚未覆盖，还需要继续判断当前数字
			x *= 2
			patches++
		}
	}
	return
}
