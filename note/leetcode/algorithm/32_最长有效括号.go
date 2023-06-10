package algorithm

// ()、(())这种理解为一个单独的括号块，判断单个括号块是否有效即可。判断成功后要注意括号块之间可能有连续，要继续判断是连续的话长度相加
func longestValidParentheses(s string) int {
	if s == "" {
		return 0
	}
	status := make([]int, len(s)) // 以当前字符结尾的字符串最长有效括号的长度 对于字符串()()) status为[0,1,0,4,0]
	// init
	status[0] = 0
	// startIndex
	max := 0
	for i := 1; i < len(s); i++ { // O(n)
		if s[i] != ')' {
			continue
		}

		if i-status[i-1]-1 < 0 { // 左侧越界
			continue
		}

		if s[i-status[i-1]-1] == '(' { // 当前字符为) 且减去以上一个字符结尾的字符的最长有效括号长度+1后的字符如果为(，则证明这个字符也是连续有效的括号
			status[i] = status[i-1] + 2
			// 当前括号块处理完毕。需要继续判断之前是否有连续的有效括号块
			if i-status[i-1]-2 >= 0 { // 每次加2之后要再往前看一位，如果status不为0，则是两个有效括号块，且连续，因此要加上。如果再往前以为的status值为0，则之前可能没有有效括号，或者有但是不连续，都不需要处理。
				if status[i-status[i-1]-2] != 0 {
					status[i] += status[i-status[i-1]-2]
				}
			}
		}

		if status[i] > max {
			max = status[i]
		}
	}
	return max
}
