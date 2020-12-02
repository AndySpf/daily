package algorithm

// 给定长度分别为 m 和 n 的两个数组，其元素由 0-9 构成，表示两个自然数各位上的数字。现在从这两个数组中选出 k (k <= m + n) 个数字拼接成一个新的数，要求从同一个数组中取出的数字保持其在原数组中的相对顺序。
//
//求满足该条件的最大数。结果返回一个表示该最大数的长度为 k 的数组。
//
//说明: 请尽可能地优化你算法的时间和空间复杂度。
//
//示例 1:
//
//输入:
//nums1 = [3, 4, 6, 5]
//nums2 = [9, 1, 2, 5, 8, 3]
//k = 5
//输出:
//[9, 8, 6, 5, 3]
// 每一个数组中找到最大子序列，那么他们按顺序拼接后肯定也是最大的。且两个最大子序列长度之和为k
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	res := make([]int, k)
	for i := 0; i <= k; i++ {
		if len(nums1) < i || len(nums2) < k-i {
			continue
		}
		sli1 := getMaxChildSli(i, nums1)
		sli2 := getMaxChildSli(k-i, nums2)
		if mergeSli := maxNumberMerge(sli1, sli2); maxNumberCompare(res, mergeSli) {
			res = mergeSli
		}
	}
	return res
}

func maxNumberMerge(sli1, sli2 []int) []int {
	merge := make([]int, len(sli1)+len(sli2))
	for i := range merge {
		if maxNumberCompare(sli1, sli2) { // 如果有相同的，则继续往下走直到第一个不相同的，看不相同的哪个大
			merge[i], sli2 = sli2[0], sli2[1:]
		} else {
			merge[i], sli1 = sli1[0], sli1[1:]
		}
	}
	return merge
}

func maxNumberCompare(sli1, sli2 []int) bool {
	for i := 0; i < len(sli1) && i < len(sli2); i++ {
		if sli1[i] != sli2[i] {
			return sli1[i] < sli2[i]
		}
	}
	return len(sli1) < len(sli2)
}

func getMaxChildSli(length int, sli []int) []int {
	if len(sli) == length {
		return sli
	}
	res := make([]int, 0, length)
	for i := range sli {
		for {
			if len(res) > 0 && sli[i] > res[len(res)-1] && len(res)+len(sli)-i > length {
				res = res[:len(res)-1]
				continue
			}
			break
		}
		if len(res) < length {
			res = append(res, sli[i])
		}
	}
	return res
}
