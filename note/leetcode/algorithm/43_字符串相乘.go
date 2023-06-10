package algorithm

import (
	"fmt"
	"strconv"
)

// 12
// 21  => 12 + 240 => 252
//-----
// 12
//240
//----
//252
// 转换为数字相乘会溢出
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1), len(num2)
	ansArr := make([]int, m+n)
	for i := m - 1; i >= 0; i-- { // 计算每一位上应该是几
		x := int(num1[i]) - '0'
		for j := n - 1; j >= 0; j-- {
			y := int(num2[j] - '0')
			ansArr[i+j+1] += x * y
		}
	}
	fmt.Println(ansArr)
	for i := m + n - 1; i > 0; i-- { // 处理进位
		ansArr[i-1] += ansArr[i] / 10
		ansArr[i] %= 10
	}
	fmt.Println(ansArr)
	ans := ""
	idx := 0
	if ansArr[0] == 0 {
		idx = 1
	}
	for ; idx < m+n; idx++ {
		ans += strconv.Itoa(ansArr[idx])
	}
	return ans
}
