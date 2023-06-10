package algorithm

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, item := range strs {
		bs := []byte(item)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		key := string(bs)
		if _, ok := groups[key]; !ok {
			groups[key] = []string{}
		}
		groups[key] = append(groups[key], item)
	}
	groupAnagramsRes := make([][]string, 0, len(groups))
	for key := range groups {
		groupAnagramsRes = append(groupAnagramsRes, groups[key])
	}
	return groupAnagramsRes
}
