package algorithm

//输入：stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
//输出：5
//解释：一种移除 5 块石头的方法如下所示：
//1. 移除石头 [2,2] ，因为它和 [2,1] 同行。
//2. 移除石头 [2,1] ，因为它和 [0,1] 同列。
//3. 移除石头 [1,2] ，因为它和 [1,0] 同行。
//4. 移除石头 [1,0] ，因为它和 [0,0] 同列。
//5. 移除石头 [0,1] ，因为它和 [0,0] 同行。
//石头 [0,0] 不能移除，因为它没有与另一块石头同行/列

func removeStones(stones [][]int) int {
	count := map[int]int{}

	var find func(x int) int
	find = func(x int) int {
		if _, ok := count[x]; !ok {
			count[x] = x
		}

		if count[x] != x {
			count[x] = find(count[x])
		}
		return count[x]
	}
	merge := func(x, y int) {
		count[find(x)] = count[find(y)]
	}

	for _, item := range stones {
		merge(item[0], item[1]+10000)
	}

	res := 0
	for i := range count {
		if count[i] == i {
			res++
		}
	}
	return len(stones) - res
}
