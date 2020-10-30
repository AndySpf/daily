package algorithm

// 自己尝试有限状态机
// "0" => true
//" 0.1 " => true
//"abc" => false
//"1 a" => false
//"2e10" => true
//" -90e3   " => true
//" 1e" => false
//"e3" => false
//" 6e-1" => true
//" 99e2.5 " => false
//"53.5e93" => true
//" --6 " => false
//"-+3" => false
//"95a54e53" => false
//
//说明: 我们有意将问题陈述地比较模糊。在实现代码之前，你应当事先思考所有可能的情况。这里给出一份可能存在于有效十进制数字中的字符列表：
//
//数字 0-9
//指数 - "e"
//正/负号 - "+"/"-"
//小数点 - "."
//当然，在输入中，这些字符的上下文也很重要。

const (
	CHAR_NUMBER_AGAIN = iota
	CHAR_E
	CHAR_SYMBOL
	CHAR_POINT_AGAIN
	CHAR_SPACE_AGAIN
	CHAR_OTHER
)

// 空格
// 正负号
// 小数点前数字
// 小数点
// 小数点后数字
// e
// e后符号
// e后整数部分
// (漏掉的末尾空格)
const (
	STATE_HEAD_SPACE = iota
	STATE_SYMBOL_BEFORE_E
	STATE_NUMBER_BEFORE_POINT
	STATE_POINT_AGAIN
	STATE_POINT_LEFT_WITHOUT_NUMBER
	STATE_NUMBER_AFTER_POINT
	STATE_E
	STATE_SYMBOL_AFTER_E
	STATE_NUMBER_AFTER_E
	STATE_TAIL_SPACE
	STATE_FAILED // 可以不用定义失败状态，每一步如果找不到往下一个状态转换的字符，就可以认为是失败状态
)

var transferState = map[int]map[int]int{
	STATE_HEAD_SPACE: {
		CHAR_SPACE_AGAIN:  STATE_HEAD_SPACE,
		CHAR_NUMBER_AGAIN: STATE_NUMBER_BEFORE_POINT,
		CHAR_POINT_AGAIN:  STATE_POINT_LEFT_WITHOUT_NUMBER,
		CHAR_SYMBOL:       STATE_SYMBOL_BEFORE_E,
	},
	STATE_SYMBOL_BEFORE_E: {
		CHAR_NUMBER_AGAIN: STATE_NUMBER_BEFORE_POINT,
		CHAR_POINT_AGAIN:  STATE_POINT_LEFT_WITHOUT_NUMBER,
		CHAR_SPACE_AGAIN:  STATE_TAIL_SPACE,
	},
	STATE_NUMBER_BEFORE_POINT: {
		CHAR_NUMBER_AGAIN: STATE_NUMBER_BEFORE_POINT,
		CHAR_POINT_AGAIN:  STATE_POINT_AGAIN,
		CHAR_SPACE_AGAIN:  STATE_TAIL_SPACE,
		CHAR_E:            STATE_E,
	},
	STATE_POINT_LEFT_WITHOUT_NUMBER: {
		CHAR_NUMBER_AGAIN: STATE_NUMBER_BEFORE_POINT,
	},
	STATE_POINT_AGAIN: {
		CHAR_NUMBER_AGAIN: STATE_NUMBER_AFTER_POINT,
		CHAR_E:            STATE_E,
		CHAR_SPACE_AGAIN:  STATE_TAIL_SPACE,
	},
	STATE_NUMBER_AFTER_POINT: {
		CHAR_NUMBER_AGAIN: STATE_NUMBER_AFTER_POINT,
		CHAR_E:            STATE_E,
		CHAR_SPACE_AGAIN:  STATE_TAIL_SPACE,
	},
	STATE_E: {
		CHAR_NUMBER_AGAIN: STATE_NUMBER_AFTER_E,
		CHAR_SYMBOL:       STATE_SYMBOL_AFTER_E,
	},
	STATE_NUMBER_AFTER_E: {
		CHAR_NUMBER_AGAIN: STATE_NUMBER_AFTER_E,
		CHAR_SPACE_AGAIN:  STATE_TAIL_SPACE,
	},
	STATE_TAIL_SPACE: {
		CHAR_SPACE_AGAIN: STATE_TAIL_SPACE,
	},
}

func byte2myChar(b byte) int {
	switch b {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER_AGAIN
	case '+', '/':
		return CHAR_SYMBOL
	case '.':
		return CHAR_POINT_AGAIN
	case 'E', 'e':
		return CHAR_E
	case ' ':
		return CHAR_SPACE_AGAIN
	default:
		return CHAR_OTHER
	}
}

func isNumberAgain(s string) bool {
	state := STATE_HEAD_SPACE // 初始化假设是空格

	for i := range s {
		if v, ok := transferState[state][byte2myChar(s[i])]; ok {
			state = v
		} else {
			return false
		}
	}
	return state != STATE_HEAD_SPACE && state != STATE_POINT_LEFT_WITHOUT_NUMBER && state != STATE_SYMBOL_BEFORE_E
}
