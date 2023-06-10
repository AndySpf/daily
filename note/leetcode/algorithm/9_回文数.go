package algorithm

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x >= 0 && x < 9 {
		return true
	}
	// 123 / 10 = 12   12 / 10 = 1
	// 123 % 10 = 3    12 % 10 = 2
	// reverse
	x1 := 0
	x2 := x
	for {
		x1 += x % 10
		x = x / 10
		if x == 0 {
			break
		}
		x1 *= 10
	}
	if x1 == x2 {
		return true
	}
	return false
}
