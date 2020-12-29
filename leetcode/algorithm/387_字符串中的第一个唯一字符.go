package algorithm

func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}
	sli := make([]int, 26)
	for i := range s {
		index := s[i] - 'a'
		if sli[index] != 0 {
			sli[index] = -1
		}
		sli[index] = i + 1
	}

	min := -1
	for index := range sli {
		if sli[index] == -1 {
			continue
		}
		if min == -1 || sli[index] < min {
			min = sli[index]
		}
	}
	return min
}
