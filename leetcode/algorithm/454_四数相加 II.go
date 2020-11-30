package algorithm

// 给定四个包含整数的数组列表 A , B , C , D ,计算有多少个元组 (i, j, k, l) ，
// 使得 A[i] + B[j] + C[k] + D[l] = 0。
// 为了使问题简单化，所有的 A, B, C, D 具有相同的长度 N，且 0 ≤ N ≤ 500 。
// 所有整数的范围在 -2**28 到 2**28 - 1 之间，最终结果不会超过 2**31 - 1 。
func fourSumCount(A []int, B []int, C []int, D []int) int {
	m1 := map[int]int{}
	for i := range A {
		for j := range B {
			if _, ok := m1[A[i]+B[j]]; !ok {
				m1[A[i]+B[j]] = 0
			}
			m1[A[i]+B[j]]++
		}
	}

	m2 := map[int]int{}
	for i := range C {
		for j := range D {
			if _, ok := m2[C[i]+D[j]]; !ok {
				m2[C[i]+D[j]] = 0
			}
			m2[C[i]+D[j]]++
		}
	}

	num := 0
	for key := range m1 {
		if _, ok := m2[0-key]; ok {
			num += m1[key] * m2[0-key]
		}
	}
	return num
}
