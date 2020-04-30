package algorithm

func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}
	if n <= 3 {
		return [][]string{}
	}
	var re [][]int

	//三个map，shus就是竖，扑面而来的爱国情怀。

	shus, pies, nas := make(map[int]bool, n), make(map[int]bool, n), make(map[int]bool, n)
	DFS := func(rows []int, n int) {}
	DFS = func(rows []int, n int) {
		row := len(rows)
		if row == n {
			aaaa := make([]int, len(rows))
			copy(aaaa, rows)
			//re = append(re,append([]int{},rows...))
			re = append(re, aaaa)
			return
		}

		for col := 0; col < n; col++ {
			if !shus[col] && !pies[row+col-1] && !nas[row-col-1] {
				shus[col] = true
				pies[row+col-1] = true
				nas[row-col-1] = true
				DFS(append(rows, col), n)
				shus[col] = false
				pies[row+col-1] = false
				nas[row-col-1] = false
			}
		}
	}

	DFS([]int{}, n)
	return bQ(re, n)
}

func bQ(re [][]int, n int) (result [][]string) {
	for _, v := range re {
		s := []string{}
		for _, vv := range v {
			str := ""
			for i := 0; i < n; i++ {
				if i == vv {
					str += "Q"
				} else {
					str += "."
				}
			}
			s = append(s, str)
		}
		result = append(result, s)
	}
	return
}
