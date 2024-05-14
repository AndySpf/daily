package algorithm

// 每次都是3，则结果可能会余1，余2
// 余数为1，则可以通过回退一次并执行两次2完成，即原数量-1+2
// 余数为2，则只用执行一次2，即原数量+1
func minimumRounds(tasks []int) int {
	m := map[int]int{}
	for i := range tasks {
		m[tasks[i]]++
	}

	res := 0
	for _, sumCount := range m {
		if sumCount == 1 { // 一定有一个数是完不成的
			return -1
		}
		res += sumCount / 3
		switch sumCount % 3 {
		case 0:
			continue
		case 1:
			res = res - 1 + 2
		case 2:
			res += 1
		}
	}
	return res
}
