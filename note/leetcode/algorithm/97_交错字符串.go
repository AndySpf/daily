package algorithm

import "fmt"

var isInterleaveRes []byte

// 回溯超时，尝试动态规划
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	isInterleaveRes = []byte{}
	return dfsIsInterleave(0, 0, s1, s2, s3)
}

func dfsIsInterleave(s1n, s2n int, s1, s2, s3 string) bool {
	fmt.Println(string(isInterleaveRes))
	if len(isInterleaveRes) == len(s3) {
		if string(isInterleaveRes) == s3 {
			return true
		}
		return false
	}

	if s1n < len(s1) {
		isInterleaveRes = append(isInterleaveRes, s1[s1n])
		if dfsIsInterleave(s1n+1, s2n, s1, s2, s3) {
			return true
		}
		isInterleaveRes = isInterleaveRes[:len(isInterleaveRes)-1]
	}

	if s2n < len(s2) {
		isInterleaveRes = append(isInterleaveRes, s2[s2n])
		if dfsIsInterleave(s1n, s2n+1, s1, s2, s3) {
			return true
		}
		isInterleaveRes = isInterleaveRes[:len(isInterleaveRes)-1]
	}
	return false
}

// abcd
// efgh
// aebfcgdh
// 动态规划
// dp[i][j] => s1前i个数和s2前j个数交错可以匹配s3的前i+j个数
// dp[i][j] = (dp[i-1][j]&&s3[i+j-1]==s1[i-1]) || (dp[i][j-1]&&s3[i+j-1]==s2[j-1])
func isInterleave1(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	// init
	dp := make([][]bool, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	dp[0][0] = true
	for i := range s1 {
		if s1[:i+1] == s3[:i+1] {
			dp[i+1][0] = true
		}
	}
	for j := range s2 {
		if s2[:j+1] == s3[:j+1] {
			dp[0][j+1] = true
		}
	}

	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			t1 := dp[i-1][j] && s3[i+j-1] == s1[i-1]
			t2 := dp[i][j-1] && s3[i+j-1] == s2[j-1]
			dp[i][j] = (t1) || (t2)
		}
	}
	return dp[len(s1)][len(s2)]
}
