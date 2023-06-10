package algorithm

import (
	"math"
)

//       0
//     6   1
//   5       2
//     4    3
//
func findRotateSteps(ring string, key string) int {
	const inf = math.MaxInt64 / 2
	n, m := len(ring), len(key)
	pos := [26][]int{}
	for i, c := range ring {
		pos[c-'a'] = append(pos[c-'a'], i)
	}
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	for _, p := range pos[key[0]-'a'] {
		dp[0][p] = minFindRotateSteps(p, n-p) + 1
	}
	for i := 1; i < m; i++ {
		for _, j := range pos[key[i]-'a'] {
			for _, k := range pos[key[i-1]-'a'] {
				dp[i][j] = minFindRotateSteps(dp[i][j], dp[i-1][k]+minFindRotateSteps(absFindRotateSteps(j-k), n-absFindRotateSteps(j-k))+1)
			}
		}
	}
	return minFindRotateSteps(dp[m-1]...)
}

func minFindRotateSteps(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if v < res {
			res = v
		}
	}
	return res
}

func absFindRotateSteps(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 超时
func findRotateSteps1(ring string, key string) int {
	m := map[byte][]int{} // ring中每个字符所在的索引
	for i := range ring {
		if _, ok := m[ring[i]]; !ok {
			m[ring[i]] = []int{}
		}
		m[ring[i]] = append(m[ring[i]], i)
	}

	m1 := make([][]int, 100)
	for i := 0; i < 100; i++ {
		m1[i] = make([]int, 100)
	}

	min := 1<<31 - 1
	var backFindRotateSteps1 func(pos int, n int, step int)
	backFindRotateSteps1 = func(pos int, n int, step int) {
		if n == len(key) {
			if step < min {
				min = step
			}
			return
		}

		item := key[n]
		for _, index := range m[item] {
			s := m1[pos][index]
			if s == 0 {
				s = transferIndex2Step(pos, len(ring), index)
			}

			//fmt.Println("refer ", pos, " targetChar ",n," index ", index, " step ", s)
			step += s
			backFindRotateSteps1(index, n+1, step)
			step -= s
		}
	}
	backFindRotateSteps1(0, 0, 0)
	return min + len(key)
}

func transferIndex2Step(refer, length, index int) int {
	index = index - refer
	if index < 0 {
		index = index + length
	}
	if index > length/2 {
		return -(index - length)
	} else {
		return index
	}
}
