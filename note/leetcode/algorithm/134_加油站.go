package algorithm

// 如果找到从x=>y油不够了，那么在[x,y)之间的任意一点到y，油都是不够的。中间这些点可以跳过检测
func canCompleteCircuit(gas []int, cost []int) int {
	var travelCircle func(begin int) (int, int)
	travelCircle = func(begin int) (int, int) {
		gasNum := gas[begin]
		costNum := cost[begin]
		pos := begin
		for {
			if gasNum < costNum {
				return -1, pos
			}

			pos++
			if pos >= len(gas) {
				pos = 0
			}

			gasNum += gas[pos]
			costNum += cost[pos]
			if pos == begin {
				break
			}
		}
		return begin, 0
	}

	for begin := 0; begin < len(gas); {
		index, end := travelCircle(begin)
		if index != -1 {
			return index
		}
		if end < begin { // 0<y<len(gas) 如果y转回切片头部，但是回不到y,那么[y:]中任意一点都到不了y即不能完成环形行驶
			return -1
		}
		begin = end + 1
	}
	return -1
}
