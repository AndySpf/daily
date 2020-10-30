package algorithm

func islandPerimeter(grid [][]int) int {
	perimeter := 0

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 0 {
				continue
			}
			perimeter += 4

			if row > 0 && grid[row-1][col] == 1 {
				perimeter -= 1
			}
			if row < len(grid)-1 && grid[row+1][col] == 1 {
				perimeter -= 1
			}
			if col > 0 && grid[row][col-1] == 1 {
				perimeter -= 1
			}
			if col < len(grid[row])-1 && grid[row][col+1] == 1 {
				perimeter -= 1
			}
		}
	}
	return perimeter
}
