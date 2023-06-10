package algorithm

import "sort"

// 名称相同且有重复邮箱，证明是同一个人
// ["join","join1@11.com"]
// ["mary","join1@11.com"]
// ["join","join1@11.com"]

//   join    mary
//
func accountsMerge(accounts [][]string) [][]string {
	m1 := map[string]int{}    // 给每个email生成一个id
	m2 := map[string]string{} // 每个email对应的账户民

	for i := range accounts {
		name := accounts[i][0]
		for _, email := range accounts[i][1:] {
			if _, ok := m1[name]; !ok {
				m1[email] = len(m1)
				m2[email] = email
			}
		}
	}

	ids := make([]int, len(m1))
	for i := range ids {
		ids[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if ids[x] != x {
			ids[x] = find(ids[x])
		}
		return ids[x]
	}

	merge := func(x, y int) {
		ids[find(x)] = find(y)
	}

	for i := range accounts {
		first := accounts[i][1]
		for _, email := range accounts[i][2:] {
			merge(m1[email], m1[first])
		}
	}

	m3 := map[int][]string{} // 每一个连通量里的email
	for email, index := range m1 {
		p := find(index)
		m3[p] = append(m3[p], email)
	}

	res := [][]string{}
	for _, emails := range m3 {
		sort.Strings(emails)
		item := append([]string{m2[emails[0]]}, emails...)
		res = append(res, item)
	}
	return res
}
