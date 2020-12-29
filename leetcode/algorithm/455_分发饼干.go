package algorithm

import "sort"

// [1,2,3]
// [1,1,1]   => 1
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	res := 0
	gPos := len(g) - 1
	for i := len(s) - 1; i >= 0; i-- {
		for j := gPos; j >= 0; j-- {
			if s[i] >= g[j] {
				res++
				gPos = j - 1
				break
			}
		}
	}
	return res
}
