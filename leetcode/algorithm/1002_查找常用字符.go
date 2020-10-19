package algorithm

func commonChars(A []string) []string {
	if len(A) == 0 {
		return nil
	}

	charMap := map[int32]int{}
	for _, char := range A[0] {
		if _, ok := charMap[char]; !ok {
			charMap[char] = 0
		}
		charMap[char]++
	}
	for _, str := range A[1:] {
		for char := range charMap {
			if count := checkChars(str, char); count == 0 {
				delete(charMap, char)
			} else if count < charMap[char] {
				charMap[char] = count
			}
		}
	}

	res := []string{}
	for char, count := range charMap {
		for i := 0; i < count; i++ {
			res = append(res, string(char))
		}
	}
	return res
}

func checkChars(str string, target int32) int {
	count := 0
	for _, char := range str {
		if char == target {
			count++
		}
	}
	return count
}
