package algorithm

//输入：s = "ADOBECODEBANC", t = "ABC"
//输出："BANC"
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	if len(s) == len(t) && s == t {
		return t
	}

	diffM := map[byte]int{}
	for i := range t {
		if _, ok := diffM[t[i]]; !ok {
			diffM[t[i]] = 0
		}
		diffM[t[i]]++
	}

	head, tail := 0, len(t)
	for i := range s[head:tail] {
		if _, ok := diffM[s[i]]; ok {
			diffM[s[i]]--
		}
	}

	if checkAllContain(diffM) {
		return s[head:tail]
	}

	minHead, minTail := 0, 1<<31-1
	for {
		if !checkAllContain(diffM) { // 当前窗口不完全包含
			if tail < len(s) {
				item := s[tail]
				if _, ok := diffM[item]; ok {
					diffM[item]--
				}
			}
			if tail == len(s) { // tail已到达末尾，此时还有不满足完全覆盖的字符串。直接跳出循环
				break
			}
			tail++
		} else {
			if tail-head == len(t) { // 已经是最小了
				return s[head:tail]
			}

			if tail-head <= minTail-minHead {
				minHead = head
				minTail = tail
			}
			// 开始收缩
			item := s[head]
			if _, ok := diffM[item]; ok {
				diffM[item]++
			}

			head++
		}

		if head >= len(s)-len(t) && tail > len(s) {
			break
		}
	}

	if minTail-minHead == 1<<31-1 { // 不存在完全包含
		return ""
	}
	return s[minHead:minTail]
}

func checkAllContain(diff map[byte]int) bool {
	for key := range diff {
		if diff[key] > 0 {
			return false
		}
	}
	return true
}
