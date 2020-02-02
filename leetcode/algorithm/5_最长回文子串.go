package algorithm

import (
	"fmt"
	"strings"
)

// abccab
// #a#b#c#c#a#b#
// 02020272020    (p[i]-1)/2
func longestPalindrome1(s string) string {
	if s == "" {
		return s
	}
	newS := fmt.Sprintf("#%s#", strings.Join(strings.Split(s, ""), "#"))
	l := make([]int, len(newS))
	max := make([]int, 2)
	for i := range newS {
		R := 0
		for {
			if i-(R+1) >= 0 && i+(R+1) < len(newS) {
				if newS[i-(R+1)] == newS[i+(R+1)] {
					R++
					continue
				}
			}
			break
		}
		if R != 0 {
			l[i] = R + 1
		}
		if l[i] > max[1] {
			max[0] = i
			max[1] = l[i]
		}
	}

	if max[0]%2 == 0 { // 偶回文数
		fmt.Println("偶回文数")
		sIndex := max[0] / 2
		offset := (max[1] - 1) / 2
		return s[sIndex-offset : sIndex+offset]
	} else {
		fmt.Println("奇回文数")
		sIndex := max[0] / 2
		offset := (max[1] - 1) / 2
		return s[sIndex-offset : sIndex+offset+1]
	}
}

type max struct {
	status []int
	size   int
}

// aba dp[0][2]
// s[i,j]是回文子串，则if dp[i-1,j+1]=dp[i,j] and (s[i+1]=s[j-1])
func longestPalindrome(s string) string {
	if s == "" {
		return s
	}
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	max := max{
		size:   0,
		status: []int{0, 0},
	}
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ { // j >= i
			// 直接判断，如果j-i < 3的情况只要s[i] == s[j]就一定是true
			if j-i == 1 || j-i == 2 || i == j {
				if s[i] == s[j] {
					dp[i][j] = true
					if j-i+1 > max.size {
						max.size = j - i + 1
						max.status[0] = i
						max.status[1] = j
					}
				}
			}
			// 直接判断结束，进入状态转换，如果dp[i-offset][j+offset]一直为回文字符串则一直往外查找
			offset := 1
			for {
				if i-offset >= 0 && j+offset < len(s) {
					if s[i-offset] == s[j+offset] {
						dp[i-offset][j+offset] = dp[i][j]
						if dp[i-offset][j+offset] { // 如果是true，则更新max，准备确定下一个状态
							if j+offset-(i-offset)+1 > max.size {
								max.size = j + offset - (i - offset) + 1
								max.status[0] = i - offset
								max.status[1] = j + offset
							}
							offset++
							continue
						}
					}
				}
				break
			}
		}
	}
	//for i := range dp {
	//	fmt.Println(dp[i])
	//}
	//fmt.Println(max)
	return s[max.status[0] : max.status[1]+1]
}
