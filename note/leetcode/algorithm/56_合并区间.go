package algorithm

import "sort"

// 输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
//输出: [[1,6],[8,10],[15,18]]
//解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	for i := 0; i < len(intervals); i++ {
		if i+1 < len(intervals) && intervals[i][1] >= intervals[i+1][0] {
			if intervals[i][0] < intervals[i+1][0] {
				intervals[i+1][0] = intervals[i][0]
			}
			if intervals[i][1] > intervals[i+1][1] {
				intervals[i+1][1] = intervals[i][1]
			}
			continue
		}
		res = append(res, intervals[i])
	}
	return res
}
