package algorithm

// ps: 时间复杂度O(n), 常熟的空间复杂度

// [1, 2, 3]  -> 4
// [-1, 0, 3] -> 1
// 哈希思路 不符合空间复杂度要求
func firstMissingPositive(nums []int) int {
	m := map[int]int{}
	for i := 1; i <= len(nums); i++ {
		m[i] = 0
	}

	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			m[nums[i]]++
		}
	}
	for i := 1; i <= len(nums); i++ {
		if m[i] == 0 {
			return i
		}
	}
	return len(nums) + 1
}

// {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
// 参考将数组看做哈希表的思路
// 0 1 2 3 4 5
// _ _ _ _ _ _
// 遍历数组，首先要明确最小正数一定是在[1,len(nums)]之间。
// 1：如果数值大于等于数组长度或者小于等于0，则右移
// 2：如果该项数值处于[1,len(nums)]之间，且值等于下标，则右移。否则
// 3：则将该值替换到对应的下表处，使目标处下标和值相同，如果要交换的位置的值和该项相同，右移。否则
// 4：对交换回来的值重复从1开始判断
// 5: 从下标为1开始遍历整理完之后的数组，如果有下标和值不同的，则返回下标，如果都相同判断第0位是否等于数组长度。不等则返回数组长度，相等则返回数组长度加1
func firstMissingPositive1(nums []int) int {
	index := 0
	for index < len(nums) {
		if nums[index] <= 0 || nums[index] >= len(nums) {
			index++
			continue
		}

		if nums[index] == index {
			index++
			continue
		}

		if nums[index] == nums[nums[index]] {
			index++
			continue
		}
		nums[index], nums[nums[index]] = nums[nums[index]], nums[index]
	}
	for i := 1; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}
	if nums[0] == len(nums) {
		return len(nums) + 1
	} else {
		return len(nums)
	}
}
