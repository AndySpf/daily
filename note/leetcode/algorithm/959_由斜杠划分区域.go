package algorithm

func regionsBySlashes(grid []string) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	nodes := make([]int, n*n*4+1) // 为了方便计算每个网格内节点的id,index=0不用
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
		nodes[find(x)] = find(y)
	}

	for i := range grid {
		for j := range grid[i] {
			ids := (i*n + j) * 4

			// 单元格内合并
			switch grid[i][j] {
			case '/':
				merge(ids+1, ids+4)
				merge(ids+2, ids+3)
			case '\\':
				merge(ids+1, ids+2)
				merge(ids+3, ids+4)
			default:
				merge(ids+1, ids+2)
				merge(ids+1, ids+3)
				merge(ids+1, ids+4)
			}

			// 单元格间合并
			if i+1 < n {
				merge(4*n+ids+1, ids+3)
			}
			if j+1 < n {
				merge(ids+8, ids+2)
			}
		}
	}

	count := -1
	for i := range nodes {
		if nodes[i] == i {
			count++
		}
	}
	return count
}
