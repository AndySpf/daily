package algorithm

// 0 0 1 1
// 1 0 1 0
// 1 1 0 0

//输入：[[0,0,1,1],[1,0,1,0],[1,1,0,0]]
//输出：39
//解释：
//转换为 [[1,1,1,1],[1,0,0,1],[1,1,1,1]]
//0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39

// 如果有一套移动方案，那么无论先执行这套方案里的哪一步，结果都是一样的
func matrixScore(A [][]int) int {
	if len(A) == 0 || len(A[0]) == 0 {
		return 0
	}
	for i := range A {
		if A[i][0] != 1 {
			for j := range A[i] {
				A[i][j] ^= 1
			}
		}
	}
	for i := 1; i < len(A[0]); i++ {
		zeroCount := 0
		for j := range A {
			if A[j][i] == 0 {
				zeroCount++
			}
			if zeroCount > len(A)/2 {
				break
			}
		}
		if zeroCount > len(A)/2 {
			for j := range A {
				A[j][i] ^= 1
			}
		}
	}

	var res int
	for i := range A {
		for j := range A[i] {
			res += A[i][j]<<len(A[i]) - j - 1
		}
	}
	return res
}

// 官方优秀题解（不需要真的翻转，只需要假设并计算最后的贡献度即可）
func matrixScore1(a [][]int) int {
	m, n := len(a), len(a[0])
	ans := 1 << (n - 1) * m // 最终结果一定第一列都是1，因此把贡献先加上
	for j := 1; j < n; j++ {
		ones := 0
		for _, row := range a {
			if row[j] == row[0] {
				ones++ // 无论第一列是什么只要和第一列相同就加1。如果是1，那么ones肯定要加1，如果是0，因为第一列不能为0，所以这一行要移动，所以这一位肯定也会变成1
			}
		}
		if ones < m-ones { // 一比零多则直接计算1的贡献度，零比一多则对这一列移动
			ones = m - ones
		}
		ans += 1 << (n - 1 - j) * ones
	}
	return ans
}
