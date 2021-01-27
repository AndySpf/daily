package algorithm

import (
	"sort"
)

func maxNumEdgesToRemove(n int, edges [][]int) int {
	if len(edges) < n-1 {
		return -1
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][0] > edges[j][0]
	})
	nodes := make([]int, n+1)
	for i := range nodes {
		nodes[i] = i
	}
	nodes1 := make([]int, len(nodes))
	c1, c2 := n, n // 连通分量的数量，初始每个节点一个连通分量

	var find func(x int) int
	find = func(x int) int {
		if nodes[x] != x {
			nodes[x] = find(nodes[x])
		}
		return nodes[x]
	}
	merge := func(x, y int) bool {
		x = find(x)
		y = find(y)
		if x == y {
			return true
		}
		nodes[x] = y
		return false
	}

	var find1 func(x int) int
	find1 = func(x int) int {
		if nodes1[x] != x {
			nodes1[x] = find1(nodes1[x])
		}
		return nodes1[x]
	}
	merge1 := func(x, y int) bool {
		x = find1(x)
		y = find1(y)
		if x == y {
			return true
		}
		nodes1[x] = y
		return false
	}

	res := 0
	cp := false
	for i := range edges {
		tp := edges[i][0]
		if tp == 3 {
			if merge(edges[i][1], edges[i][2]) {
				res++
			} else {
				c1--
				c2--
			}
		} else if tp == 2 {
			if !cp {
				copy(nodes1, nodes)
				cp = true
			}
			if merge1(edges[i][1], edges[i][2]) {
				res++
			} else {
				c2--
			}
		} else if tp == 1 {
			if merge(edges[i][1], edges[i][2]) {
				res++
			} else {
				c1--
			}
		}
	}
	if c1 > 1 || c2 > 1 {
		return -1
	}
	return res
}
