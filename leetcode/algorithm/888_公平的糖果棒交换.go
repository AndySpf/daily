package algorithm

func fairCandySwap(A []int, B []int) []int {
	sumA, sumB := 0, 0
	for i := range A {
		sumA += A[i]
	}
	m1 := map[int]int{}
	for i := range B {
		sumB += B[i]
		m1[B[i]]++
	}

	diff := (sumA - sumB) / 2
	for i := range A {
		if _, ok := m1[A[i]-diff]; ok {
			return []int{A[i], A[i] - diff}
		}
	}
	return nil
}
