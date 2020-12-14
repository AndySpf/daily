package algorithm

// 更新，插入，删除
// dp[i][j]  word1的前i个字符到word2的前j个字符需要经过的步骤
// dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1]+1)
func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		return len(word1) + len(word2)
	}
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}

	// 目标： hrose => ros， 可分为下列三种情况，取最小
	// hors, r  dp[i][j-1]次操作后hros已经变为r，则对于B的前j位，只需要在A中合适位置插入一个B[j]就好了，加一
	// hor, ro  dp[i-1][j]次操作后hro已经变为ro 则对于A的前i位，只需要在B中添加A[i]就好了(B中添加等效于A中删除)，加一
	// hor, r   dp[i-1][j-1],则A将第i位更新为B的第j位即可，如果word2[j]!=word1[i],则加1
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			min := dp[i-1][j-1]
			if word2[j-1] != word1[i-1] {
				min++
			}
			if min > dp[i-1][j]+1 {
				min = dp[i-1][j] + 1
			}
			if min > dp[i][j-1]+1 {
				min = dp[i][j-1] + 1
			}
			dp[i][j] = min
		}
	}
	return dp[len(word1)][len(word2)]
}

// leetcode提供的便于理解的思路：
// 问题1：如果 word1[0..i-1] 到 word2[0..j-1] 的变换需要消耗 k 步，那 word1[0..i] 到 word2[0..j] 的变换需要几步呢？
//
//答：先使用 k 步，把 word1[0..i-1] 变换到 word2[0..j-1]，消耗 k 步。再把 word1[i] 改成 word2[j]，就行了。如果 word1[i] == word2[j]，什么也不用做，一共消耗 k 步，否则需要修改，一共消耗 k + 1 步。
//
//问题2：如果 word1[0..i-1] 到 word2[0..j] 的变换需要消耗 k 步，那 word1[0..i] 到 word2[0..j] 的变换需要消耗几步呢？
//
//答：先经过 k 步，把 word1[0..i-1] 变换到 word2[0..j]，消耗掉 k 步，再把 word1[i] 删除，这样，word1[0..i] 就完全变成了 word2[0..j] 了。一共 k + 1 步。
//
//问题3：如果 word1[0..i] 到 word2[0..j-1] 的变换需要消耗 k 步，那 word1[0..i] 到 word2[0..j] 的变换需要消耗几步呢？
//
//答：先经过 k 步，把 word1[0..i] 变换成 word2[0..j-1]，消耗掉 k 步，接下来，再插入一个字符 word2[j], word1[0..i] 就完全变成了 word2[0..j] 了。
//
//从上面三个问题来看，word1[0..i] 变换成 word2[0..j] 主要有三种手段，用哪个消耗少，就用哪个。
