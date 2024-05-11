package algorithm

// 输入：garbage = ["G","P","GP","GG"], travel = [2,4,3] => 21
func garbageCollection(garbage []string, travel []int) int {
	total := len(garbage[0])
	pos := map[byte]int{
		'M': 0,
		'P': 0,
		'G': 0,
	}
	var sumTravel = func(from, to int) (sum int) {
		for _, v := range travel[from:to] {
			sum += v
		}
		return
	}
	for i := 1; i < len(garbage); i++ {
		for _, b := range []byte(garbage[i]) {
			// 如果一个垃圾车当前位置小于index，则往前移动n个距离并更新位置
			if pos[b] < i {
				total += sumTravel(pos[b], i)
				pos[b] = i
			}
			// 收拾垃圾的时间
			total++
		}
	}
	return total
}
