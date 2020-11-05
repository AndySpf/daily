package algorithm

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if len(needle) == len(haystack) {
		if needle == haystack {
			return 0
		} else {
			return -1
		}
	}
	if needle == "" {
		return 0
	}

	l := len(needle)
	for i := 0; i <= len(haystack)-l; i++ {
		if haystack[i] == needle[0] && haystack[i:i+l] == needle {
			return i
		}
	}
	return -1
}
