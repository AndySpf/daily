package algorithm

// S = "ab#c", T = "ad#c" => true
// S = "ab#e#c", T = "ad#c" => true
// S = "ab##", T = "c#d#" => true
func backspaceCompare(S string, T string) bool {
	stackS := make([]byte, 0, len(S))
	for i := 0; i < len(S); i++ {
		if S[i] == '#' {
			if len(stackS) > 0 {
				stackS = stackS[:len(stackS)-1]
			}
		} else {
			stackS = append(stackS, S[i])
		}
	}

	stackT := make([]byte, 0, len(T))
	for i := 0; i < len(T); i++ {
		if T[i] == '#' {
			if len(stackT) > 0 {
				stackT = stackT[:len(stackT)-1]
			}
		} else {
			stackT = append(stackT, T[i])
		}
	}
	s := string(stackS)
	t := string(stackT)
	if s == t {
		return true
	}
	return false
}

// 想到反向遍历但没想好要怎么对比
// 查看题解：针对每个字符串，需要有一个表示忽略个数的变量，遇到#时加1，遇到其他字符时判断是否大于0，大于0则减1，直到第一个为0且不是#的字符，然后双方做对比
