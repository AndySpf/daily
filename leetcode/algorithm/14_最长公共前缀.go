package algorithm

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	max := strs[0]
	for i := 1; i < len(strs); i++ {
		max = checkPrefix(max, strs[i])
	}
	return max
}

func checkPrefix(pre string, s string) string {
	for i := 0; i < len(pre); i++ {
		if i == len(s) {
			return s
		}

		if pre[i] != s[i] {
			return pre[:i]
		}
	}
	return pre
}
