package algorithm

// ()[]{}
func isValid(s string) bool {
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []byte{}
	for i := range s {
		if s[i] == ' ' {
			continue
		}
		if _, ok := m[s[i]]; ok {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || m[stack[len(stack)-1]] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
