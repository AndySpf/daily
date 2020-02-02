package algorithm

func mapEqual(m1, m2 map[string]uint8) bool {
	if len(m1) != len(m2) {
		return false
	}
	for m1Key := range m1 {
		if _, ok := m2[m1Key]; !ok {
			return false
		}
		if m1[m1Key] != m2[m1Key] {
			return false
		}
	}
	return true
}

func counter(data []string) map[string]uint8 {
	res := map[string]uint8{}
	for i := range data {
		if _, ok := res[data[i]]; ok {
			res[data[i]]++
		} else {
			res[data[i]] = 1
		}
	}
	return res
}

func splitStringByStep(s string, step int) []string {
	if len(s)%step != 0 {
		return []string{}
	}
	num := len(s) / step
	res := make([]string, num)
	for i := 0; i < num; i++ {
		res[i] = s[i*step : (i+1)*step]
	}
	return res
}

//  s = "barfothefoobarman",words = ["foo","bar"]  -> [0, 9]
func findSubstring(s string, words []string) []int {
	if len(words) == 0 || s == "" {
		return nil
	}
	wordLen := len(words[0])
	if len(s) < len(words)*wordLen {
		return nil
	}

	result := []int{}
	wordsCounter := counter(words)
	for i := 0; i <= len(s)-len(words)*wordLen; i++ {
		window := s[i : i+len(words)*wordLen]
		windowCounter := counter(splitStringByStep(window, wordLen))
		if mapEqual(wordsCounter, windowCounter) {
			result = append(result, i)
		}
	}
	return result
}

// 太暴力了。。极端用例超时
//  s = "bar foo the foo bar man",words = ["foo","bar"]  -> [0, 9]
//  s = "wordgoodgoodgoodbest",words = ["word","good","best"]  -> []
func findSubstring1(s string, words []string) []int {
	if s == "" {
		return []int{}
	}

	step := len(words[0])

	result := []int{}

	for i := range s {
		tmpI := i

		tmpWords := make([]string, len(words))
		copy(tmpWords, words)
		for {
			if tmpI+step > len(s) {
				break
			}
			exist, index := inSlice(tmpWords, s[tmpI:tmpI+step])
			if !exist { // 不存在直接结束
				break
			}

			copy(tmpWords[index:], tmpWords[index+1:])
			tmpWords = tmpWords[:len(tmpWords)-1]
			if len(tmpWords) == 0 && exist { // 存在且tmpWords长度为0了则证明全部被匹配到了
				result = append(result, i)
			}
			tmpI += step // 剩下的存在但是还没匹配完的情况则继续匹配
		}
	}
	return result
}

func inSlice(slice []string, s string) (bool, int) {
	for i := range slice {
		if s == slice[i] {
			return true, i
		}
	}
	return false, -1
}
