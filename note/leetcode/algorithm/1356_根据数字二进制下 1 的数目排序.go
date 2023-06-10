package algorithm

import "sort"

// 十进制转为二进制后有多少个1
// 7 111
// 6 110
// 5 101

// 10000
// 1 4 5 3 7

type sortByBitsStruct struct {
	nums []int
	m    map[int]int
}

func (s sortByBitsStruct) Len() int {
	return len(s.nums)
}

func (s sortByBitsStruct) Less(i, j int) bool {
	if s.m[s.nums[i]] < s.m[s.nums[j]] {
		return true
	} else if s.m[s.nums[i]] == s.m[s.nums[j]] {
		if s.nums[i] < s.nums[j] {
			return true
		}
	}
	return false
}

func (s sortByBitsStruct) Swap(i, j int) {
	s.nums[i], s.nums[j] = s.nums[j], s.nums[i]
}

func sortByBits(arr []int) []int {
	s := sortByBitsStruct{
		nums: arr,
		m:    map[int]int{},
	}
	for i := range arr {
		s.m[arr[i]] = get1Count(arr[i])
	}
	sort.Sort(s)
	return s.nums
}

func get1Count(num int) (count int) {
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return
}
