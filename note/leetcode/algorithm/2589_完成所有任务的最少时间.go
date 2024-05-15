package algorithm

import (
	"sort"
)

// 输入：tasks = [[2,3,1],[4,5,1],[1,5,2]]
// 输出：2

// 输入：tasks = [[1,3,2],[2,5,3],[5,6,2]]
// 输出：4

// 参考贪心解法
func findMinimumTime(tasks [][]int) int {
	if len(tasks) == 0 {
		return 0
	}
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i][1] == tasks[j][1] {
			return tasks[i][2] > tasks[j][2]
		}
		return tasks[i][1] < tasks[j][1]
	})
	// [[1,3,2],[2,5,3],[5,6,2]]
	var minimumTime int
	nums := make([]int, tasks[len(tasks)-1][1]+1)
	findExecTime := func(start, end int) (total int) {
		for i := start; i <= end; i++ {
			if nums[i] == 1 {
				total++
			}
		}
		return
	}
	for i := range tasks {
		total := findExecTime(tasks[i][0], tasks[i][1])
		if tasks[i][2] <= total {
			continue
		} else {
			diff := tasks[i][2] - total
			index := tasks[i][1]
			for diff > 0 {
				if nums[index] != 1 {
					nums[index] = 1
					diff--
					minimumTime++
				}
				index--
			}
		}
	}
	return minimumTime
}
