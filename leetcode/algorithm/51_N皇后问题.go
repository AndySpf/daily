package algorithm

import "fmt"

// 位运算代替数组存储状态
func solveNQueens1(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}
	if n <= 3 {
		return nil
	}

	count = n
	result = [][]string{}

	queues := []int{}
	status := map[string]int{}
	status["col"] = 0
	status["tiltLeft"] = 0
	status["tiltRight"] = 0
	back1(status, 0, queues)
	return result
}

func back1(status map[string]int, row int, queues []int) {
	if row > count-1 {
		if len(queues) == count {
			fmt.Println(queues)
			result = append(result, generateNQueue(queues))
		}
		return
	}

	for i := 0; i < count; i++ {
		col := i % count
		left := row + col
		right := row + (count - 1 - col)
		if (status["col"]>>col&1) == 0 && (status["tiltLeft"]>>left&1) == 0 && (status["tiltRight"]>>right&1) == 0 {
			status["col"] |= 1 << col
			status["tiltLeft"] |= 1 << left
			status["tiltRight"] |= 1 << right
			queues = append(queues, row*count+i)
			back1(status, row+1, queues)
			queues = queues[:len(queues)-1]
			status["col"] ^= 1 << col
			status["tiltLeft"] ^= 1 << left
			status["tiltRight"] ^= 1 << right
		}
	}
	return
}

/*
. Q . .
. . . Q
Q . . .
. . Q .
*/

var count int
var result [][]string

func solveNQueens(n int) [][]string {
	if n == 1 {
		return [][]string{{"Q"}}
	}
	if n <= 3 {
		return nil
	}

	count = n
	result = [][]string{}

	queues := []int{}
	status := map[string][]bool{}
	status["col"] = make([]bool, n)
	status["tiltLeft"] = make([]bool, 2*n-1)
	status["tiltRight"] = make([]bool, 2*n-1)
	back(status, 0, queues)
	return result
}

func back(status map[string][]bool, row int, queues []int) {
	if row > count-1 {
		if len(queues) == count {
			fmt.Println(queues)
			result = append(result, generateNQueue(queues))
		}
		return
	}

	for i := 0; i < count; i++ {
		col := i % count
		left := row + col
		right := row + (count - 1 - col)
		if !status["col"][col] && !status["tiltLeft"][left] && !status["tiltRight"][right] {
			status["col"][col] = true
			status["tiltLeft"][left] = true
			status["tiltRight"][right] = true
			queues = append(queues, row*count+i)
			back(status, row+1, queues)
			queues = queues[:len(queues)-1]
			status["col"][col] = false
			status["tiltLeft"][left] = false
			status["tiltRight"][right] = false
		}
	}
	return
}

func generateNQueue(source []int) []string {
	tmp := make([]string, count)
	for i, pos := range source {
		p := pos % count
		for j := 0; j < count; j++ {
			if j != p {
				tmp[i] += "."
			} else {
				tmp[i] += "Q"
			}
		}
	}
	return tmp
}
