package algorithm

func combinationSum(candidates []int, target int) [][]int {
	targetCombinationSum = target
	allCombinationSum = candidates
	resultCombinationSum = [][]int{}

	backCombinationSum(0, []int{})
	return resultCombinationSum
}

func backCombinationSum(index int, item []int) {
	v := sumInt(item)
	if v >= targetCombinationSum {
		if v == targetCombinationSum {
			var newItem = make([]int, len(item))
			copy(newItem, item)
			resultCombinationSum = append(resultCombinationSum, newItem)
		}
		return
	}

	for i := index; i < len(allCombinationSum); i++ {
		if allCombinationSum[i] > targetCombinationSum {
			continue
		}
		item = append(item, allCombinationSum[i])
		backCombinationSum(i, item)
		item = item[:len(item)-1]
	}
}
