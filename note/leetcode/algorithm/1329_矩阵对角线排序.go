package algorithm

import "sort"

// 输入：mat = [[3,3,1,1],[2,2,1,2],[1,1,1,2]]
// 输出：[[1,1,1,1],[1,2,2,2],[1,2,3,3]]
// 3 3 1 1
// 2 2 1 2
// 1 1 1 2
// 当前思路是转换出一个完整的transfer，其实可以每一次循环中直接就排序然后重新赋值，能减少遍历次数
// 或者原地冒泡排序
func diagonalSort(mat [][]int) [][]int {
	if len(mat) <= 1 {
		return mat
	}
	if len(mat[0]) <= 1 {
		return mat
	}

	transfer := make([][]int, len(mat)+len(mat[0])-1)
	index := 0

	for startCol, startRow := 0, len(mat)-1; startCol <= len(mat[0])-1; {
		var col, row = startCol, startRow
		transfer[index] = make([]int, 0)
		for {
			transfer[index] = append(transfer[index], mat[row][col])
			row++
			col++
			if row >= len(mat) || col >= len(mat[0]) {
				break
			}
		}
		index++

		if startRow == 0 {
			startCol++
		} else {
			startRow--
		}
	}

	for i := range transfer {
		sort.Ints(transfer[i])
	}

	index = 0
	for startCol, startRow := 0, len(mat)-1; startCol <= len(mat[0])-1; {
		var col, row = startCol, startRow
		for i := range transfer[index] {
			mat[row][col] = transfer[index][i]
			row++
			col++
		}
		index++

		if startRow == 0 {
			startCol++
		} else {
			startRow--
		}
	}

	return mat
}
