package algorithm

// 给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	direction := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	direIndex := 0
	posX, posY := 0, 0
	res[posX][posY] = 1
	for i := 2; i <= n*n; {
		posX += direction[direIndex][0]
		posY += direction[direIndex][1]
		// 两种条件转弯：最外圈超限；内圈碰到已经被赋值过的
		if posX < 0 || posY < 0 || posX == n || posY == n || res[posX][posY] != 0 {
			posX -= direction[direIndex][0]
			posY -= direction[direIndex][1]
			direIndex++
			if direIndex >= len(direction) {
				direIndex = 0
			}
			continue
		}
		res[posX][posY] = i
		i++
	}
	return res
}
