package algorithm

// x ** 10 => 1010 => x**2 * x**8
// x ** 77 => 1001101 => x ** 64 * x**8 * x**4 * x
// 参考官方奇妙的思路
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, N int) float64 {
	ans := 1.0
	// 贡献的初始值为 x
	x_contribute := x
	// 在对 N 进行二进制拆分的同时计算答案
	for N > 0 {
		if N%2 == 1 {
			// 如果 N 二进制表示的最低位为 1，那么需要计入贡献
			ans *= x_contribute
		}
		// 将贡献不断地平方
		x_contribute *= x_contribute
		// 舍弃 N 二进制表示的最低位，这样我们每次只要判断最低位即可
		N /= 2
	}
	return ans
}

func myPow1(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1.0 / quickMul(x, -n)
}

func quickMul1(x float64, N int) float64 {
	contribution := x
	ans := 1.0

	for N > 0 {
		if N%2 == 1 { // 最低位为1
			ans *= contribution
		}
		contribution *= contribution
		N /= 2 // 每次右移一位
	}
	return ans
}
