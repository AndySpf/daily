package algorithm

func findRedundantConnection(edges [][]int) []int {
	points := make([]int, len(edges)+1)
	for i := range points {
		points[i] = i
	}

	var find func(x int) int
	find = func(x int) int {
		if points[x] != x {
			points[x] = find(points[x])
		}
		return points[x]
	}

	merge := func(x, y int) bool {
		x = find(x)
		y = find(y)
		points[x] = points[y]
		if x == y {
			return true
		}
		return false
	}

	for _, item := range edges {
		if merge(item[0], item[1]) {
			return item
		}
	}
	return nil
}
