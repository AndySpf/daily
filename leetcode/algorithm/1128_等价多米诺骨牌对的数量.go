package algorithm

func numEquivDominoPairs(dominoes [][]int) int {
	m := map[int]int{}
	for _, domino := range dominoes {
		if domino[0] > domino[1] {
			domino[0], domino[1] = domino[1], domino[0]
		}
		k := domino[0]*10 + domino[1]
		m[k]++
	}
	count := 0
	for _, c := range m {
		if c == 1 {
			continue
		}
		count += (c*c - 1) / 2
	}
	// c*c-1*c-2*c-3...*2
	// 2*c-2*c-3...*2
	return count
}
