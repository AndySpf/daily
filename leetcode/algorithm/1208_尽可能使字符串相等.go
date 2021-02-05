package algorithm

// a,b,c,d  b,c,d,e
func equalSubstring(s string, t string, maxCost int) int {
	left, right, cost := 0, 0, getCost(t[0], s[0])
	for right < len(s) {
		if cost > maxCost {
			cost -= getCost(s[left], t[left])
			left++
		}

		right++
		if right < len(s) {
			cost += getCost(s[right], t[right])
		}
	}
	return right - left
}

func getCost(u1, u2 uint8) int {
	if u1 > u2 {
		return int(u1 - u2)
	}
	return int(u2 - u1)
}
