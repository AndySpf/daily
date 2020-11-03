package algorithm

import (
	"strings"
)

//输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]
func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	res := []string{}

	var backGenerateParenthesis func(s []string, left, right int)

	backGenerateParenthesis = func(s []string, left, right int) {
		if len(s) == 2*n {
			res = append(res, strings.Join(s, ""))
			return
		}
		if left < n {
			s = append(s, "(")
			backGenerateParenthesis(s, left+1, right)
			s = s[:len(s)-1]
		}
		if right < left {
			s = append(s, ")")
			backGenerateParenthesis(s, left, right+1)
			s = s[:len(s)-1]
		}
	}
	backGenerateParenthesis([]string{}, 0, 0)
	return res
}

func generateParenthesis1(n int) []string {
	dp := map[int][]string{}
	dp[0] = []string{""}
	dp[1] = []string{"()"}
	dp[2] = []string{"()()", "(())"}
	for i := 3; i <= n; i++ {
		for p := 0; p <= i-1; p++ {
			q := i - 1 - p
			for j := range dp[q] {
				for k := range dp[p] {
					// dp[i]认为相对于dp[i-1]在最左边加了左括号，然后右括号只可能存在于dp[i-1]的中间或者右边
					// 因此dp[i]就等于"(" + dp[p][k]+ ")" 的组合加上dp[q][j]的组合，q+p=i-1
					dp[i] = append(dp[i], "("+dp[p][k]+")"+dp[q][j])
				}
			}

		}
	}
	return dp[n]
}
