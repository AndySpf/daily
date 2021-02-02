package algorithm

//输入：s = "AABABBA", k = 1
//输出：4
//解释：
//将中间的一个'A'替换为'B',字符串变为 "AABBBBA"。
//子串 "BBBB" 有最长重复字母, 答案为 4。
// dp[i][j] 前i个字符中替换j个
func characterReplacement(s string, k int) int {
	if len(s) < 2 {
		return len(s)
	}
	left, right := 0, 0
	nums := make([]int, 26)
	nums[s[0]-'A']++
	maxIndex := s[0] - 'A'
	for right < len(s) { // 判断当前窗口是否可以变成相同字符串(窗口内出现次数最多的字符+k>=窗口宽度)
		if nums[maxIndex]+k < right-left+1 {
			nums[s[left]-'A']--
			left++
		}

		// 如果到边界，则右边不移动
		right++
		if right < len(s) {
			nums[s[right]-'A']++
			if nums[s[right]-'A'] > nums[maxIndex] {
				maxIndex = s[right] - 'A'
			}
		}
	}
	return right - left
}
