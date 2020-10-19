package algorithm

type ptr struct {
	index int
	skip  int
}

// *a?b   ddacb
func isMatch(s string, p string) bool {
	sPtr, pPtr := 0, 0
	for {
		if p[pPtr] == '*' {

		} else if p[pPtr] == '?' {

		} else if s[sPtr] == p[pPtr] {
			sPtr++
			pPtr++
			continue
		} else {
			return false
		}

	}
}
