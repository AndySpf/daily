package algorithm

func addToArrayForm(A []int, K int) []int {
	if K == 0 {
		return A
	}
	k := []int{}
	for K != 0 {
		k = append(k, K%10)
		K = K / 10
	}
	upper := 0
	for p1, p2 := 0, len(A)-1; p1 >= len(k) && p2 < 0; {
		add := upper
		upper = 0
		if p1 < len(k) {
			add += k[p1]
		}
		if p2 > 0 {
			add += A[p2]
		}
		if add > 9 {
			upper = 1
			add = add % 10
		}

		if p2 > 0 {
			A[p2] = add
		} else {
			A = append([]int{add}, A...)
		}
		p2--
		p1++
	}
	if upper == 1 {
		A = append([]int{1}, A...)
	}
	return A
}
