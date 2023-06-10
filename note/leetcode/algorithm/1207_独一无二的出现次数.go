package algorithm

func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}
	for i := range arr {
		if _, ok := m[i]; !ok {
			m[i] = 0
		}
		m[i]++
	}

	m1 := map[int]int{}
	for num, count := range m {
		if _, ok := m1[count]; ok {
			return false
		}
		m1[count] = num
	}
	return true
}
