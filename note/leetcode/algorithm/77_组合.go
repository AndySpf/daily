package algorithm

// 组合数：C(m,n) = A(m,n) / n!; A(m,n) = m! / (m-n)! => C(m,n) = m! / n!(m-n)!
var all, width int
var res [][]int

func combine(n int, k int) [][]int {
	if n < k {
		return nil
	}

	if n == k {
		t := make([]int, n)
		for i := 1; i <= n; i++ {
			t[i-1] = i
		}
		return [][]int{t}
	}

	all = n
	width = k
	res = [][]int{}

	backCombine(1, make([]int, 0, k))
	return res
}

func backCombine(index int, item []int) {
	if len(item) == width {
		newItem := make([]int, len(item))
		copy(newItem, item)
		res = append(res, newItem)
		return
	}
	for i := index; i <= all; i++ {
		item = append(item, i)
		backCombine(i+1, item)
		item = item[:len(item)-1]
	}
	return
}
