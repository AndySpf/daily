package algorithm

// 优先给顾客10元的零钱
func lemonadeChange(bills []int) bool {
	m := map[int]int{}
	for _, money := range bills {
		switch money {
		case 10:
			if m[5] <= 0 {
				return false
			}
			m[5]--
		case 20:
			if m[10] >= 1 && m[5] >= 1 {
				m[10]--
				m[5]--
			} else if m[5] >= 3 {
				m[5] -= 3
			} else {
				return false
			}
		}
		m[money]++
	}
	return true
}
