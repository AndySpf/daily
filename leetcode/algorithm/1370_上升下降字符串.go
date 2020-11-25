package algorithm

import "fmt"

// 输入：s = "aaaabbbbcccc"
// 输出："abccbaabccba"
func sortString(s string) string {
	charArr := make([]int, 26)
	for i := range s {
		charArr[s[i]-'a']++
	}

	result := ""
	pos := 0
	symbol := 1
	for {
		allEmpty := true
		fmt.Println(charArr)
		for {
			if charArr[pos] != 0 {
				allEmpty = false
				result += string('a' + pos)
				charArr[pos] -= 1
			}
			pos += symbol
			if pos > len(charArr)-1 || pos < 0 {
				pos -= symbol
				break
			}
		}
		symbol = -symbol
		if allEmpty {
			break
		}
	}
	return result
}
