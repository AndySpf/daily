package algorithm

//beginWord = "hit",
//endWord = "cog",
//wordList = ["hot","dot","dog","lot","log","cog"]
//
//输出: 5
//
//解释: 一个最短转换序列是 "hit" -> "hot" -> "dot" -> "dog" -> "cog",
//     返回它的长度 5。
// 时间复杂度很高，应该学习下官方的优化建图
func ladderLength(beginWord string, endWord string, wordList []string) int {
	m := map[string][]string{
		beginWord: []string{},
	}
	for i := range wordList {
		tmp := []string{}
		for key := range m {
			if validLadderLength(key, wordList[i]) {
				m[key] = append(m[key], wordList[i])
				tmp = append(tmp, key)
			}
		}
		m[wordList[i]] = tmp
	}
	if _, ok := m[endWord]; !ok {
		return 0
	}

	// 根据连通性广度优先搜索
	steps := make(map[string]int, len(wordList)+1)
	for i := range wordList {
		steps[wordList[i]] = 1<<32 - 1
	}
	steps[beginWord] = 1
	queue := []string{beginWord}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if item == endWord {
			return steps[item]
		}

		for _, next := range m[item] {
			if steps[next] == 1<<32-1 {
				queue = append(queue, next)
				steps[next] = steps[item] + 1
			}
		}
	}
	return 0
}

func validLadderLength(s1 string, s2 string) bool {
	count := 0 // 差异度
	for i := range s1 {
		if s1[i] != s2[i] {
			count++
		}
		if count > 1 {
			return false
		}
	}
	return true
}
