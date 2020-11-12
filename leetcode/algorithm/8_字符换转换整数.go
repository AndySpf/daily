package algorithm

const (
	HEAD_EMPTY = iota
	FIRST_NONEMPTY_SYMBOL
	FIRST_NONEMPTY_NUMBER
	FIRST_NONEMPTY_OTHER
	NUMBER_CONTINUE
	NUMBER_CONTINUE_BROKEN
)

func myAtoi(s string) int {
	max := int64((1 << 31) - 1)
	//minFindRotateSteps := int64(-(1 << 31))
	if len(s) == 0 {
		return 0
	}
	s1 := searchNumbers(s)
	if len(s1) <= 0 {
		return 0
	}

	symbol := 1
	if s1[0] == '-' || s1[0] == '+' {
		if s1[0] == '-' {
			symbol = ^symbol + 1
		}
		s1 = s1[1:]
	}

	res := int64(0)
	for index, bt := range []byte(s1) {
		res += int64(bt) - 48
		if index != len(s1)-1 {
			res *= 10
		}
		if symbol < 0 && res >= max+1 {
			res = max + 1
		}
		if res >= max {
			res = max
		}
	}
	if symbol < 0 {
		res = -res
	}

	return int(res)
}

// 可以优化下， 类似第7题里面。状态机找到的字符串直接组结果
func searchNumbers(s string) string {
	state := HEAD_EMPTY
	res := ""
	for i := range s {
		if s[i] == ' ' && state == HEAD_EMPTY {
			state = HEAD_EMPTY
			continue
		}
		if (s[i] == '+' || s[i] == '-') && state == HEAD_EMPTY {
			res += string(s[i])
			state = FIRST_NONEMPTY_SYMBOL
			continue
		}
		if s[i] >= '0' && s[i] <= '9' && state == HEAD_EMPTY {
			res += string(s[i])
			state = FIRST_NONEMPTY_NUMBER
			continue
		}
		if s[i] >= '0' && s[i] <= '9' && state == FIRST_NONEMPTY_SYMBOL {
			res += string(s[i])
			state = NUMBER_CONTINUE
			continue
		}
		if s[i] >= '0' && s[i] <= '9' && state == FIRST_NONEMPTY_NUMBER {
			res += string(s[i])
			state = NUMBER_CONTINUE
			continue
		}
		if s[i] >= '0' && s[i] <= '9' && state == NUMBER_CONTINUE {
			res += string(s[i])
			state = NUMBER_CONTINUE
			continue
		}
		state = FIRST_NONEMPTY_OTHER
		break
	}
	if len(res) == 1 && (res[0] == '-' || res[0] == '+') {
		return ""
	}

	return res
}
