package algorithm

import "fmt"

//输入：isConnected = [[1,1,0],[1,1,0],[0,0,1]]
//输出：2

func findCircleNum(isConnected [][]int) int {
	if len(isConnected) < 2 {
		return len(isConnected)
	}
	steps := make([]int, len(isConnected))
	provinceCount := 0
	for i := range steps {
		if steps[i] == 0 {
			queue := []int{i}
			for len(queue) != 0 {
				item := queue[0]
				queue = queue[1:]
				steps[item] = 1
				for i := range isConnected[item] {
					if steps[i] == 1 || isConnected[item][i] == 0 {
						continue
					}
					queue = append(queue, i)
				}
			}
			provinceCount++
		}
	}
	return provinceCount
}

func findCircleNum1(isConnected [][]int) (ans int) {
	n := len(isConnected)
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(from, to int) {
		parent[find(from)] = find(to)
	}

	for i, row := range isConnected {
		for j := i + 1; j < n; j++ {
			if row[j] == 1 {
				union(i, j)
			}
		}
	}
	fmt.Println(parent)
	for i, p := range parent {
		if i == p {
			ans++
		}
	}
	return
}
