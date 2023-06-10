package algorithm

import (
	"sort"
)

// 给你一个points 数组，表示 2D 平面上的一些点，其中 points[i] = [xi, yi] 。
//连接点 [xi, yi] 和点 [xj, yj] 的费用为它们之间的
// 曼哈顿距离 ：|xi - xj| + |yi - yj| ，其中 |val| 表示 val 的绝对值。
//请你返回将所有点连接的最小总费用。只有任意两点之间 有且仅有 一条简单路径时，才认为所有点都已连接。
func minCostConnectPoints(points [][]int) int {
	// 计算任意两点间距离判断最小生成树?prime算法和Kruskal算法
	if len(points) < 1 {
		return 0
	}
	type edge struct {
		edge int
		p1   int
		p2   int
	}

	getDistance := func(x1, y1, x2, y2 int) int {
		tmp1, tmp2 := x1-x2, y1-y2
		if tmp1 < 0 {
			tmp1 = -tmp1
		}
		if tmp2 < 0 {
			tmp2 = -tmp2
		}
		return tmp1 + tmp2
	}

	edges := make([]edge, 0, (len(points)*(len(points)-1))/2)
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			edges = append(edges, edge{
				edge: getDistance(points[i][0], points[i][1], points[j][0], points[j][1]),
				p1:   i,
				p2:   j,
			})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].edge < edges[j].edge
	})

	ps := make([]int, len(points))
	for i := range ps {
		ps[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if ps[x] != x {
			ps[x] = find(ps[x])
		}
		return ps[x]
	}

	merge := func(x, y int) bool {
		x = find(x)
		y = find(y)
		if x == y {
			return true
		}
		ps[x] = y
		return false
	}

	res := 0
	for i := range edges {
		if !merge(edges[i].p1, edges[i].p2) {
			res += edges[i].edge
		}
	}
	return res
}
