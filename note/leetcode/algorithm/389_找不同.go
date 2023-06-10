package algorithm

func findTheDifference(s string, t string) byte {
	if len(s) == 0 {
		return t[0]
	}

	m1 := map[byte]int{}
	for i := range t {
		m1[t[i]]++
	}

	for i := range s {
		m1[s[i]]--
		if m1[s[i]] == 0 {
			delete(m1, s[i])
		}
	}
	for key := range m1 {
		return key
	}
	return ' '
}

func findTheDifference1(s string, t string) byte {
	sli := make([]int, 26)
	for i := range s {
		sli[s[i]-'a']++
		sli[t[i]-'a']--
	}
	sli[t[len(t)-1]-'a']--
	for i := range sli {
		if sli[i] != 0 {
			return byte(i + 'a')
		}
	}
	return ' '
}
