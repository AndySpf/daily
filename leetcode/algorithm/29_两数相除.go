package algorithm

// 10 / 2 2+2+2+2+2
// 10 / -2 2+2+2+2+2
func divide(dividend int, divisor int) int {
	if dividend == 0 || divisor == 0 {
		return 0
	}

	res := 0
	symbol := int8(1)
	// 确定符号
	if dividend < 0 {
		dividend = ^dividend + 1
		symbol = ^symbol + 1
	}
	if divisor < 0 {
		divisor = ^divisor + 1
		symbol = ^symbol + 1
	}

	// 剪枝
	if dividend < divisor {
		return 0
	}
	if divisor == 1 { // 越界只会是除数绝对值为1时出现如： -1<<31 / -1 == 1<<31 溢出
		res = dividend
		if (symbol < 0 && res > 1<<31) || (symbol > 0 && res > 1<<31-1) {
			res = 1<<31 - 1
		}
	} else {
		// 计算除法
		step := 0
		for { // 20 / 2   14  6
			dividend = dividend - (divisor << step)
			if dividend >= 0 {
				res += 1 << step
				step++
				if dividend == 0 {
					break
				}
			} else if step > 0 {
				dividend += divisor << step
				step--
			} else {
				break
			}
		}
	}

	if symbol < 0 {
		res = ^res + 1
	}
	return res
}
