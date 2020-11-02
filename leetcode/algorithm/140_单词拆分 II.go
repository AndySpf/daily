package algorithm

import (
	"strings"
)

var wordBreakRes = []string{}

// "aaabaaaaa" ["a","aa","aaaa"]
// "bcadef" ["bc", "afed", "bca", "def"]
// "pineapplepenapple", ["apple","pen","applepen","pine","pineapple"]
// 回溯超时
func wordBreak(s string, wordDict []string) []string {
	if len(s) == 0 || len(wordDict) == 0 {
		return nil
	}
	wordBreakRes = []string{}
	m := map[string]int{}
	for i := range wordDict {
		m[wordDict[i]] = i
	}
	wordBreakRollback(s, m, 0, []int{}, wordDict)
	return wordBreakRes
}

func wordBreakRollback(s string, m map[string]int, start int, queue []int, sli []string) {
	if start == len(s) {
		tmp := make([]string, len(queue))
		for i := range queue {
			tmp[i] = sli[queue[i]]
		}
		wordBreakRes = append(wordBreakRes, strings.Join(tmp, " "))
		return
	}
	for i := start; i < len(s); i++ {
		if index, ok := m[s[start:i+1]]; ok {
			queue = append(queue, index)
			wordBreakRollback(s, m, i+1, queue, sli)
			queue = queue[:len(queue)-1]
		}
	}
	return
}

// 记忆化搜索，记录s中索引为i开始，能组成哪些单词的列表
func wordBreak1(s string, wordDict []string) (sentences []string) {
	wordSet := map[string]struct{}{}
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}

	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(index int) [][]string
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		wordsList := [][]string{}
		for i := index + 1; i <= n; i++ {
			word := s[index:i]
			if _, has := wordSet[word]; has {
				if i == n {
					wordsList = append(wordsList, []string{word})
				} else {
					for _, nextWords := range backtrack(i) {
						wordsList = append(wordsList, append([]string{word}, nextWords...))
					}
				}
			}
		}

		dp[index] = wordsList
		return wordsList
	}
	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}
	return
}
