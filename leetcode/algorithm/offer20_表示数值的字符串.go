package algorithm

import (
	"strings"
)

// 字符串首位必须为:数字/+/-/.
// .后面必须为 数字/e/空, 且.只能出现一次
// +,-如果不在开头，那就必须要在e/E后出现,且后面必须为 数字/.
// e/E前面必须为数字且后面也必须为+/-/数字,且出现e/E后不能有.或者e/E
func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return false
	}
	onePoint := false
	for i, item := range s {
		if i == 0 {
			if item >= 48 && item <= 57 {
				continue
			} else if (item == '+' || item == '-') && len(s) > 1 {
				continue
			} else if item == '.' && len(s) > 1 && s[1] >= 48 && s[1] <= 57 {
				onePoint = true
				continue
			} else {
				return false
			}
		}

		if i > 0 {
			switch item {
			case '.':
				if onePoint {
					return false
				}
				if i == len(s)-1 {
					if s[i-1] >= 48 && s[i-1] <= 57 {
						onePoint = true
						continue
					} else {
						return false
					}
				}
				if ((s[i+1] >= 48 && s[i+1] <= 57) || (s[i+1] == 'e' || s[i+1] == 'E')) && ((s[i-1] >= 48 && s[i-1] <= 57) || (s[i-1] == '+' || s[i-1] == '-')) {
					onePoint = true
					continue
				} else {
					return false
				}

			case '+', '-':
				if i >= len(s)-1 {
					return false
				}
				if ((s[i+1] >= 48 && s[i+1] <= 57) || s[i+1] == '.') && (s[i-1] == 'e' || s[i-1] == 'E') {
					continue
				} else {
					return false
				}
			case 'e', 'E':
				if i >= len(s)-1 {
					return false
				}
				if ((s[i+1] >= 48 && s[i+1] <= 57) || s[i+1] == '+' || s[i+1] == '-') && ((s[i-1] >= 48 && s[i-1] <= 57) || s[i-1] == '.') {
					for index := i + 1; index < len(s); index++ {
						if s[index] == '.' || s[index] == 'e' || s[index] == 'E' {
							return false
						}
					}
					continue
				} else {
					return false
				}
			default:
				if item >= 48 && item <= 57 {
					continue
				}
				return false
			}
		}
	}
	return true
}

// 有限状态机
// " +1.1e-1"
//1. 空格
//2. 运算符
//3. 整数位
//4. 左侧有整数的小数点
//5. 左侧无整数的小数点
//6. 小数部分
//7. 字符e
//8. 指数部分的符号位
//9. 指数部分的整数部分
//10. 末尾的空格
// 则有效状态为3，4，6，9，10
// 状态真难找正确【官方给的技巧：那么怎么挖掘出所有可能的状态呢？一个常用的技巧是，用「当前处理到字符串的哪个部分」当作状态的表述】
type status int
type valueType int

const (
	STATE_INIT status = iota
	STATE_INT_SIGN
	STATE_INTEGER
	STATE_POINT
	STATE_POINT_LEFT_WITHOUT_INT
	STATE_FRACTION
	STATE_EXP
	STATE_EXP_SIGN
	STATE_EXP_INTEGER
	STATE_END
)

const (
	CHAR_NUMBER valueType = iota
	CHAR_EXP
	CHAR_POINT
	CHAR_SIGN
	CHAR_SPACE
	CHAR_ILLEGAL
)

var transferMap = map[status]map[valueType]status{
	STATE_INIT: {
		CHAR_SPACE:  STATE_INIT,
		CHAR_NUMBER: STATE_INTEGER,
		CHAR_SIGN:   STATE_INT_SIGN,
		CHAR_POINT:  STATE_POINT_LEFT_WITHOUT_INT,
	},
	STATE_INT_SIGN: {
		CHAR_NUMBER: STATE_INTEGER,
		CHAR_POINT:  STATE_POINT_LEFT_WITHOUT_INT,
	},
	STATE_INTEGER: {
		CHAR_NUMBER: STATE_INTEGER,
		CHAR_EXP:    STATE_EXP,
		CHAR_POINT:  STATE_POINT,
		CHAR_SPACE:  STATE_END,
	},
	STATE_POINT: {
		CHAR_NUMBER: STATE_FRACTION,
		CHAR_EXP:    STATE_EXP,
		CHAR_SPACE:  STATE_END,
	},
	STATE_POINT_LEFT_WITHOUT_INT: {
		CHAR_NUMBER: STATE_FRACTION,
	},
	STATE_FRACTION: {
		CHAR_NUMBER: STATE_FRACTION,
		CHAR_EXP:    STATE_EXP,
		CHAR_SPACE:  STATE_END,
	},
	STATE_EXP: {
		CHAR_NUMBER: STATE_EXP_INTEGER,
		CHAR_SIGN:   STATE_EXP_SIGN,
	},
	STATE_EXP_SIGN: {
		CHAR_NUMBER: STATE_EXP_INTEGER,
	},
	STATE_EXP_INTEGER: {
		CHAR_NUMBER: STATE_EXP_INTEGER,
		CHAR_SPACE:  STATE_END,
	},
	STATE_END: {
		CHAR_SPACE: STATE_END,
	},
}

func toValueType(ch byte) valueType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER
	case 'e', 'E':
		return CHAR_EXP
	case '.':
		return CHAR_POINT
	case '+', '-':
		return CHAR_SIGN
	case ' ':
		return CHAR_SPACE
	default:
		return CHAR_ILLEGAL
	}
}

func isNumber1(s string) bool {
	state := STATE_INIT
	for i := 0; i < len(s); i++ {
		typ := toValueType(s[i])
		if _, ok := transferMap[state][typ]; !ok {
			return false
		} else {
			state = transferMap[state][typ]
		}
	}
	return state == STATE_INTEGER || state == STATE_POINT || state == STATE_FRACTION || state == STATE_EXP_INTEGER || state == STATE_END
}
