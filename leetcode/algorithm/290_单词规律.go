package algorithm

import "strings"

// 给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。
// 这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串
// str 中的每个非空单词之间存在着双向连接的对应规律。
func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	m1 := map[byte][]int{}
	m2 := map[string][]int{}
	for i := range pattern {
		m1[pattern[i]] = append(m1[pattern[i]], i)
		m2[words[i]] = append(m2[words[i]], i)
	}
	if len(m1) != len(m2) {
		return false
	}

	for _, indexs := range m1 {
		if len(indexs) <= 1 {
			continue
		}
		word := words[indexs[0]]
		for _, index := range indexs[1:] {
			if words[index] != word {
				return false
			}
		}
	}
	return true
}

func wordPattern1(pattern string, s string) bool {
	m1 := map[string]byte{} // 一个字符串对应一个byte
	m2 := map[byte]string{} // 一个byte对应一个字符串
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	for i := range pattern {
		if _, ok := m1[words[i]]; !ok {
			m1[words[i]] = pattern[i]
			continue
		}
		if m1[words[i]] != pattern[i] {
			return false
		}

		if _, ok := m2[pattern[i]]; !ok {
			m2[pattern[i]] = words[i]
			continue
		}
		if m2[pattern[i]] != words[i] {
			return false
		}
	}
	return true
}
