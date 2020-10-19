package algorithm

var nQueens2status map[string][]bool
var nQueens2count int

func totalNQueens(n int) int {
	if n == 1 {
		return 1
	}
	if n < 4 {
		return 0
	}

	nQueens2status = map[string][]bool{}
	nQueens2count = 0

	nQueens2status["cols"] = make([]bool, n)
	nQueens2status["left"] = make([]bool, n*2)
	nQueens2status["right"] = make([]bool, n*2)
	for i := 0; i < n; i++ {
		nQueens2status["cols"][i] = false
		nQueens2status["left"][i] = false
		nQueens2status["right"][i] = false
	}

	nQueens(n, 0)
	return nQueens2count
}

// 0,7 -> 0    0,0->0
// 1,7 -> 1    0,1->1
// 1,6 -> 2    0,2->2
// 2,7 -> 2    1,3->4
// 2,5 -> 4	   2,1->3
func nQueens(n, row int) {
	if row == n {
		nQueens2count++
		return
	}

	for i := 0; i < n; i++ {
		if !checkNQueens(n, row, i) {
			continue
		}

		nQueens2status["cols"][i] = true
		nQueens2status["left"][i+row] = true
		nQueens2status["right"][n-1-i+row] = true

		nQueens(n, row+1)

		nQueens2status["cols"][i] = false
		nQueens2status["left"][i+row] = false
		nQueens2status["right"][n-1-i+row] = false
	}
}

func checkNQueens(n, row, col int) bool {
	if !nQueens2status["cols"][col] &&
		!nQueens2status["left"][col+row] &&
		!nQueens2status["right"][n-1-col+row] {
		return true
	}
	return false
}
