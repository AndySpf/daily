package algorithm

import "fmt"

//输入:
//[
// [ 1, 2, 3 ],
// [ 4, 5, 6 ],
// [ 7, 8, 9 ]
//]
//输出: [1,2,3,6,9,8,7,4,5]

var direction = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func spiralOrder(matrix [][]int) []int {
	layers, flag := (len(matrix)+1)/2, len(matrix)%2
	if (len(matrix[0])+1)/2 < layers {
		layers = (len(matrix[0]) + 1) / 2
		flag = len(matrix[0]) % 2
	}

	res := make([]int, 0, len(matrix)*len(matrix[0]))
	for layer := 0; layer < layers; layer++ {
		pos := []int{layer, layer}
		res = append(res, matrix[layer][layer])
		for i := range direction {
			// 最后一层可能会只剩下一条边，如果是则走一半方向即可
			if layer == layers-1 && flag == 1 && i >= 2 {
				break
			}
			for {
				if x, y, ok := move(matrix, pos, direction[i], layer); ok {
					fmt.Println(x, "|", y, "|", direction[i])
					res = append(res, matrix[y][x])
					pos[0], pos[1] = x, y
					continue
				}
				break
			}
		}
	}
	fmt.Println(res)
	return res
}

func move(sli [][]int, last, next []int, layer int) (int, int, bool) {
	weight, height := len(sli[0])-layer, len(sli)-layer
	x1 := next[0] + last[0]
	y1 := next[1] + last[1]

	if x1 >= weight || x1 < layer || y1 < layer || y1 >= height || (x1 == layer && y1 == layer) {
		return 0, 0, false
	} else {
		return x1, y1, true
	}
}
