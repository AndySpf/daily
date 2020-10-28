package algorithm

import "strings"

// 2=>b,c,d
// 3=>e,f,g
// 4=>h,i,j

var letterCombinationsRes []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	letterCombinationsRes = []string{}
	queue := []string{}
	dfsLetterCombinations(m, digits, 0, queue)
	return letterCombinationsRes
}

func dfsLetterCombinations(m map[byte][]string, s string, n int, queue []string) {
	if n == len(s) {
		str := strings.Join(queue, "")
		letterCombinationsRes = append(letterCombinationsRes, str)
		return
	}

	for _, item := range m[s[n]] {
		queue = append(queue, item)
		dfsLetterCombinations(m, s, n+1, queue)
		queue = queue[:len(queue)-1]
	}
}
