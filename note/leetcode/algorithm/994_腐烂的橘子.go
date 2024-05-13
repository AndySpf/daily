package algorithm

// 2,1,1
// 0,1,1
// 1,0,1
func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	find := func(i, j int) int {
		var changeCount int
		for _, delta := range [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}} {
			deltaX := i + delta[0]
			deltaY := j + delta[1]
			if (deltaX >= 0 && deltaX < m) && (deltaY >= 0 && deltaY < n) && grid[deltaX][deltaY] == 1 {
				changeCount++
				grid[deltaX][deltaY] = 2
			}
		}
		grid[i][j] = 3
		return changeCount
	}
	res := 0
	for {
		changeCount := 0
		fresh := 0
		source := [][2]int{}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] == 1 {
					fresh++
				}
				if grid[i][j] == 2 {
					source = append(source, [2]int{i, j})
				}
			}
		}
		for _, v := range source {
			changeCount += find(v[0], v[1])
		}
		if changeCount == 0 {
			if fresh > 0 {
				return -1
			}
			break
		}
		res++
	}
	return res
}
