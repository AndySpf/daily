package algorithm

// 给定一个非负整数数组，你最初位于数组的第一个位置。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个位置。

// 贪心, 从某一位置开始字符值减去后面字符索引与当前索引的差值，再减去后面那个字符值大于0则跳过，小于0，则更新位置。
// [2,3,1,1,4]  true
// [3,2,1,0,4]  false
func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	cur := 0
	for i := 1; i < len(nums); i++ {
		if nums[cur] == 0 {
			return false
		}

		if nums[cur]-(i-cur)-nums[i] > 0 {
			continue
		}

		cur = i
		if nums[cur]+i >= len(nums) {
			return true
		}
	}
	return true
}
