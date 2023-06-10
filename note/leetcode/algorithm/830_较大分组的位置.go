package algorithm

// "abbxxxxzzy"
// [3,6]
func largeGroupPositions(s string) [][]int {
	if len(s) < 3 {
		return nil
	}
	slow, fast := 0, 1
	res := [][]int{}
	for fast < len(s) {
		if s[fast] == s[slow] {
			fast++
			continue
		}
		if fast-1-slow >= 2 {
			res = append(res, []int{slow, fast - 1})
		}
		slow, fast = fast, fast+1
	}
	return res
}
