package algorithm

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	start, end := 0, len(A)-1
	for {
		if start == len(A)-1 || end == 0 {
			return false
		}

		if A[start] < A[start+1] {
			start++
		} else if A[end] < A[end-1] {
			end--
		} else {
			break
		}
	}

	return start == end
}
