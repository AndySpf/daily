package algorithm

var (
	targetCombinationSum int
	allCombinationSum    []int
	resultCombinationSum [][]int
	allCombinationSumS   []combinationSum2S
)

type combinationSum2S struct {
	Num   int
	Count int
}

func combinationSum2(candidates []int, target int) [][]int {
	targetCombinationSum = target
	allCombinationSumS = combinationSumDump(candidates)
	resultCombinationSum = [][]int{}

	backCombinationSum2(0, []int{})
	return resultCombinationSum
}

func backCombinationSum2(index int, item []int) {
	v := sumInt(item)
	if v >= targetCombinationSum {
		if v == targetCombinationSum {
			var newItem = make([]int, len(item))
			copy(newItem, item)
			resultCombinationSum = append(resultCombinationSum, newItem)
		}
		return
	}

	for i := index; i < len(allCombinationSumS); i++ {
		if allCombinationSumS[i].Num > targetCombinationSum {
			continue
		}
		item = append(item, allCombinationSumS[i].Num)
		allCombinationSumS[i].Count--
		if allCombinationSumS[i].Count == 0 {
			backCombinationSum2(i+1, item)
		} else {
			backCombinationSum2(i, item)
		}
		allCombinationSumS[i].Count++
		item = item[:len(item)-1]
	}
}

func sumInt(data []int) int {
	var res int
	for i := range data {
		res += data[i]
	}
	return res
}

func combinationSumDump(source []int) []combinationSum2S {
	m := map[int]int{}
	for i := range source {
		if _, ok := m[source[i]]; ok {
			m[source[i]]++
			continue
		}
		m[source[i]] = 1
	}
	var res = make([]combinationSum2S, 0, len(m))
	for key, value := range m {
		res = append(res, combinationSum2S{
			Num:   key,
			Count: value,
		})
	}
	return res
}
