package algorithm

// 123 -> 321
func reverseInt(x int) int {
	if x > -10 && x < 10 {
		return x
	}

	x1 := 0
	for {
		x1 += x % 10
		x = x / 10
		if x == 0 {
			break
		}
		x1 *= 10
	}

	if x1 > (1<<31-1) || x1 < -(1<<31-1) {
		return 0
	}

	return x1
}
