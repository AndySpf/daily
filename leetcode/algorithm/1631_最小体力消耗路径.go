package algorithm

import (
	"sort"
)

// heights = [[1,2,2],[3,8,2],[5,3,5]]
// 如果限制了走的方向就可以用动态规划了
// 这个可以用并查集，控制第0个点到最后一个点之间的连通性。height由低到高添加合并
func minimumEffortPath(heights [][]int) int {
	type edge struct {
		n1, n2 int
		height int
	}

	edges := []edge{}
	for i := range heights {
		for j := range heights[0] {
			if i < len(heights)-1 {
				edges = append(edges, edge{
					n1:     i*len(heights[0]) + j,
					n2:     (i+1)*len(heights[0]) + j,
					height: abs(heights[i][j] - heights[i+1][j]),
				})
			}
			if j < len(heights[0])-1 {
				edges = append(edges, edge{
					n1:     i*len(heights[0]) + j,
					n2:     i*len(heights[0]) + j + 1,
					height: abs(heights[i][j] - heights[i][j+1]),
				})
			}
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].height < edges[j].height
	})

	nodes := make([]int, len(heights)*len(heights[0]))
	for i := range nodes {
		nodes[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if nodes[x] != x {
			nodes[x] = find(nodes[x])
		}
		return nodes[x]
	}
	merge := func(x, y int) {
		x = find(x)
		y = find(y)
		nodes[x] = y
	}
	for i, item := range edges {
		merge(item.n1, item.n2)
		if find(0) == find(nodes[len(nodes)-1]) {
			return edges[i].height
		}
	}
	return 0
}
