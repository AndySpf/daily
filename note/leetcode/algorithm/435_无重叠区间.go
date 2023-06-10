package algorithm

import "sort"

//输入: [ [1,2], [2,3], [3,4], [1,3] ]
//输出: 1
//解释: 移除 [1,3] 后，剩下的区间没有重叠。
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		//不能按照区间起始位置排，如果有[1,4] [2,3]这种，按照起始位置排，就排错了
		return intervals[i][1] < intervals[j][1]
	})
	count := 0
	pos := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= pos {
			pos = intervals[i][1]
		} else {
			count++
		}
	}
	return count
}
