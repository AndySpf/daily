package algorithm

import (
	"sort"
)

//输入：points = [[3,9],[7,12],[3,8],[6,8],[9,10],[2,9],[0,9],[3,9],[0,6],[2,8]]
//输出：2
// 右边界最靠左的一项必须耗费一支箭（按照每一项中右边界排升序），否则这一项没有箭可以打掉，
// 选定这个右边界后，凡是左边界小于等于他的都可以被引爆。如果大于则说明这支箭能发挥的最大作用已经没了。
// 此时忽略已经引爆掉的气球。相当于从头开始，将这一项的右边界更新为最左右边界
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
	maxRight := points[0][1]
	ans := 1
	for _, p := range points {
		if p[0] > maxRight {
			maxRight = p[1]
			ans++
		}
	}
	return ans
}
