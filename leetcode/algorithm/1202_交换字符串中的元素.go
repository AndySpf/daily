package algorithm

import (
	"sort"
)

//输入：s = "dcab", pairs = [[0,3],[1,2],[0,2]]
//输出："abcd"
//解释：
//交换 s[0] 和 s[3], s = "bcad"
//交换 s[0] 和 s[2], s = "acbd"
//交换 s[1] 和 s[2], s = "abcd"

func smallestStringWithSwaps(s string, pairs [][]int) string {
	top := make([]int, len(s))
	for i := range s {
		top[i] = i
	}

	m := map[int][]byte{}

	var find func(x int) int
	find = func(x int) int {
		if top[x] != x {
			top[x] = find(top[x])
		}
		return top[x]
	}

	merge := func(i, j int) {
		i, j = find(i), find(j)
		top[i] = j
	}

	for _, item := range pairs {
		merge(item[0], item[1])
	}
	for i := range top {
		m[find(top[i])] = append(m[find(top[i])], s[i])
	}
	for _, v := range m {
		sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
	}

	ans := make([]byte, len(s))
	for i := range ans {
		f := find(i)
		ans[i] = m[f][0]
		m[f] = m[f][1:]
	}
	return string(ans)
}
