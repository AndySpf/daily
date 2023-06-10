package algorithm

import "sort"

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	allCells := make([][]int, 0, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			allCells = append(allCells, []int{i, j})
		}
	}

	sort.Slice(allCells, func(i, j int) bool {
		a := abs(allCells[i][0]-r0) + abs(allCells[i][1]-c0)
		b := abs(allCells[j][0]-r0) + abs(allCells[j][1]-c0)
		return a < b
	})
	return allCells
}
