package algorithm

func sortArrayByParityII(A []int) []int {
	for i := range A {
		if i%2 == A[i]%2 {
			continue
		}
		for j := i + 1; j < len(A); j++ {
			if i%2 == A[j]%2 {
				A[i], A[j] = A[j], A[i]
				break
			}
		}
	}
	return A
}

func sortArrayByParityII1(A []int) []int {
	i, j := 0, 1
	for j < len(A) {
		if i%2 == 0 {
			i += 2
		} else {
			if j%2 == 1 {
				j += 2
			} else {
				A[i], A[j] = A[j], A[i]
			}
		}

	}
	return A
}
