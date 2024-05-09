package algorithm

import "container/heap"

type secondElementItem struct {
	index int
	value int
}

type secondElement []secondElementItem

func (s secondElement) Len() int {
	return len(s)
}

func (s secondElement) Less(i, j int) bool {
	return s[i].value < s[j].value
}

func (s secondElement) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *secondElement) Push(x any) {
	*s = append(*s, x.(secondElementItem))
}

func (s *secondElement) Pop() any {
	a := *s
	b := a[len(a)-1]
	*s = a[:len(a)-1]
	return b
}

// [2,4,0,9,6] -> [9,6,6,-1,-1]
// 两个小顶堆，分别代表还有两条命、一条命的数。每来一个新的数就尝试从两个堆中杀掉一条命
func secondGreaterElement(nums []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}
	second := new(secondElement)
	one := new(secondElement)
	for i := 0; i < len(nums); i++ {
		for {
			if one.Len() > 0 {
				v := heap.Pop(one).(secondElementItem)
				if nums[i] > v.value {
					res[v.index] = nums[i]
					continue
				}
				heap.Push(one, v)
			}
			break
		}

		for {
			if second.Len() > 0 {
				v := heap.Pop(second).(secondElementItem)
				if nums[i] > v.value {
					heap.Push(one, v)
					continue
				}
				heap.Push(second, v)
			}
			break
		}
		heap.Push(second, secondElementItem{index: i, value: nums[i]})
	}
	return res
}
