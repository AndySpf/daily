package algorithm

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	for i := range nums1 {
		exist := false
		for j := range nums2 {
			if nums2[j] == nums1[i] {
				exist = true
				break
			}
		}

		if exist {
			add := true
			for k := range res {
				if res[k] == nums1[i] {
					add = false
				}
			}
			if add {
				res = append(res, nums1[i])
			}
		}
	}
	return res
}
