package algorithm

// d*a?b   ddasdacb
func isMatch(s string, p string) bool {
	status := make([][]bool, len(s)+1)
	for i := range status {
		status[i] = make([]bool, len(p)+1)
	}

	status[0][0] = true
	for j := 0; j < len(p); j++ {
		if p[j] == '*' && status[0][j] {
			status[0][j+1] = true
		} else {
			break
		}
	}
	// $abcd
	// $ab*cd
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if p[j] == '?' {
				status[i+1][j+1] = status[i][j]
			}

			if p[j] == '*' {
				status[i+1][j+1] = status[i+1][j] || status[i][j+1] // 使用*号则取决于status[i][j+1](ab能匹配到ab*,则abc一定也能)， 不使用*号则取决于status[i+1][j](ab能匹配到ab,则ab也一定能匹配到ab*)
			}

			if p[j] == s[i] {
				status[i+1][j+1] = status[i][j]
			}
		}
	}
	return status[len(s)][len(p)]
}
