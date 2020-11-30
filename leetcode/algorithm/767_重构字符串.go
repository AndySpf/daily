package algorithm

import (
	"sort"
)

// 示例 1:
//输入: S = "aab"
//输出: "aba"
//
//示例 2:
//输入: S = "aaab"
//输出: ""
type reorganizeItem struct {
	char  byte
	count int
}

// 当且仅当总长度为奇数，最多字符数量为(len(S)+1)/2时。才必须要求最多字符全部排在奇数位
// 为了照顾这种情况，按次数排序后，从次数多到次数少的字符，先从0索引下标开始安排字符。走到底之后从1再走一次
// 两轮下来就可以把所有字符安排妥当（要考虑清楚的是：某一个字符先走0，然后又走1 会不会相邻）
func reorganizeString(S string) string {
	if len(S) < 2 {
		return S
	}
	sli := make([]reorganizeItem, 26)
	for i := range S {
		sli[S[i]-'a'].char = S[i]
		sli[S[i]-'a'].count++
		if sli[S[i]-'a'].count > (len(S)+1)/2 {
			return ""
		}
	}
	sort.Slice(sli, func(i, j int) bool {
		return sli[i].count < sli[j].count
	})
	res := make([]byte, len(S))
	index := 0
	for i := len(sli) - 1; i >= 0; i-- {
		if sli[i].count == 0 {
			break
		}
		for j := 0; j < sli[i].count; j++ {
			res[index] = sli[i].char
			index += 2
			if index > len(res)-1 {
				index = 1
			}
		}
	}
	return string(res)
}
