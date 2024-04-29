package algorithm

// 正数进制转换
// 2^1 + 2^0 = 3 -> 11
// 2^2 = 4 -> 100
// 2^2 + 2^1 + 2^0 = 7 -> 111
// 2 ^ 3 + 2^1 =10 -> 1010 2*2*2*2 10/2 / 2 / 2 / 2   0101
// 2 ^ 3 + 2^0 + 2^1 = 11 -> 1011 2*2*2*2 11/2 / 2 / 2 / 2 1101

// 负数进制转换
// (-2)^2 + (-2)^1 = 2 -> 110             2/-2  /-2   01
// (-2)^2 + (-2)^1 + (-2)^ 0 = 3 -> 111   3/-2 / -2   11
// (-2)^2 = 4 -> 100					  4/-2 / -2 / -2 001
// (-2 ^ 4 + -2^3) + (-2^2 + -2^1) + (-2^0)= 11 -> 11111  2*2*2*2 11/2 / 2 / 2 / 2 1101

// 2 ^ 2 + 2 ^ 1 = 110
// -2^4 + -2^3 + -2 ^ 1= "11010"   / 6 / -2 / -2        -3，0 ｜ 2，1 ｜ -1， 0 ｜ 1， 1 ｜ 0，1
func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	return baseTransfer(n, -2)
}

// 普通二进制转换
func baseTransfer(n int, base int) string {
	res := []byte{}
	for n != 0 {
		remainder := n % base
		n = n / base
		if remainder < 0 {
			n += 1
			remainder = remainder - base
		}
		res = append(res, '0'+byte(remainder))
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i] = res[len(res)-i], res[i]
	}
	return string(res)
}
