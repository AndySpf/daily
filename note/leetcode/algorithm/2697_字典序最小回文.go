package algorithm

func makeSmallestPalindrome(s string) string {
	n := len(s)
	if n == 1 {
		return s
	}
	bytes := []byte(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		if s[i] == s[j] {
			continue
		}
		if s[i] > s[j] {
			bytes[i] = bytes[j]
		} else {
			bytes[j] = bytes[i]
		}
	}
	return string(bytes)
}
