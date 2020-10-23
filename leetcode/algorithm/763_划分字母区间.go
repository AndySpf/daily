package algorithm

//输入：S = "ababcbacadefegdehijhklij"
//输出：[9,7,8]
//划分结果为 "ababcbaca", "defegde", "hijhklij"。
func partitionLabels(S string) (partitionLabelsResult []int) {
	if S == "" {
		return nil
	}
	if len(S) == 1 {
		return []int{1}
	}

	lastIndex := map[byte]int{}
	for i := range S {
		lastIndex[S[i]] = i
	}

	start, end := 0, lastIndex[S[0]]
	for i := range S {
		if lastIndex[S[i]] > end {
			end = lastIndex[S[i]]
		}
		if i == end {
			partitionLabelsResult = append(partitionLabelsResult, end-start+1)
			start = end + 1
		}
	}
	return partitionLabelsResult
}

// 官方题解用26长度的顺序表代替哈希表
func partitionLabels1(s string) (partition []int) {
	lastPos := [26]int{}
	for i, c := range s {
		lastPos[c-'a'] = i
	}
	start, end := 0, 0
	for i, c := range s {
		if lastPos[c-'a'] > end {
			end = lastPos[c-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return
}
