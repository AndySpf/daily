package algorithm

import "sort"

// [-4,-1,0,3,10]
func sortedSquares(A []int) []int {
	if len(A) == 0 {
		return nil
	}
	result := make([]int, len(A))
	head, end, index := 0, len(A)-1, len(A)-1
	for {
		headTmp := A[head] * A[head]
		endTmp := A[end] * A[end]
		if head == end {
			result[index] = A[head] * A[head]
			break
		}
		if headTmp >= endTmp {
			result[index] = headTmp
			head++
		}
		if headTmp < endTmp {
			result[index] = endTmp
			end--
		}
		index--
	}
	return result
}

func sortedSquares1(A []int) []int {
	for i := range A {
		A[i] = A[i] * A[i]
	}
	sort.Ints(A)
	return A
}
