package algorithm

var (
	combinationSum3Range int
)

func combinationSum3(k int, n int) [][]int {
	combinationSum3Range = k
	targetCombinationSum = n
	resultCombinationSum = [][]int{}

	backCombinationSum3(1, []int{})
	return resultCombinationSum
}

func backCombinationSum3(num int, item []int) {
	v := sumInt(item)
	if v == targetCombinationSum && len(item) == combinationSum3Range {
		var newItem = make([]int, combinationSum3Range)
		copy(newItem, item)
		resultCombinationSum = append(resultCombinationSum, newItem)
	}
	if v > targetCombinationSum {
		return
	}
	if len(item) > combinationSum3Range {
		return
	}

	for i := num; i <= 9; i++ {
		if i > targetCombinationSum {
			continue
		}
		item = append(item, i)
		backCombinationSum3(i+1, item)
		item = item[:len(item)-1]
	}
}
