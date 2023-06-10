package algorithm

import "sort"

// [1,2,3]
// [1,1,1]   => 1
// 假设最多饼干满足胃口最大的人，第二多饼干满足胃口第二大的人。
// 如果最多的饼干给胃口第二大的人吃，那么第二多的饼干就满足不了胃口最大的人。不符合满足最多人
// 因此只要保证胃口最大的人吃最多的饼干即可
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
