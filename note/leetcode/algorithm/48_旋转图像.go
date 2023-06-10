package algorithm

//给定 matrix =
//[
//  [1,2,3],
//  [4,5,6],
//  [7,8,9]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [7,4,1],
//  [8,5,2],
//  [9,6,3]
//]
func rotate(matrix [][]int) {
	n := len(matrix)
	if n == 1 {
		return
	}
	layers := n / 2
	for layer := 0; layer < layers; layer++ {
		for i := 0; i < (len(matrix)-layer*2)-1; i++ {
			//a, b := matrix[layer][layer+i], matrix[layer+i][n-1-layer]
			//c, d := matrix[n-1-layer-i][layer], matrix[n-1-layer][n-1-layer-i]
			//fmt.Println(a,"|",b,"|",d,"|",c)
			matrix[layer][layer+i], matrix[layer+i][n-1-layer], matrix[n-1-layer][n-1-layer-i], matrix[n-1-layer-i][layer] = matrix[n-1-layer-i][layer], matrix[layer][layer+i], matrix[layer+i][n-1-layer], matrix[n-1-layer][n-1-layer-i]
		}
	}
}
