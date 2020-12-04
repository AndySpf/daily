package algorithm

import (
	"container/heap"
	"sort"
)

//示例 2：
//
//输入: [1,2,3,3,4,4,5,5]
//输出: True
//解释:
//你可以分割出这样两个连续子序列 :
//1, 2, 3, 4, 5
//3, 4, 5

func isPossible(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	countM := map[int]int{}
	for i := range nums {
		countM[nums[i]]++
	}

	m := map[int]int{} // 以x结尾的子序列数量
	for _, v := range nums {
		if countM[v] == 0 { // 因为创建子序列可能某些数已经被消耗完了，不能再使用他们了
			continue
		}
		if m[v-1] > 0 {
			m[v-1]--
			m[v]++
			countM[v]--
		} else {
			if countM[v+1] > 0 && countM[v+2] > 0 {
				m[v+2]++
				countM[v]--
				countM[v+1]--
				countM[v+2]--
			} else {
				return false
			}
		}
	}
	return true
}

type hpIsPossible struct{ sort.IntSlice }

func (h *hpIsPossible) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hpIsPossible) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hpIsPossible) push(v int) { heap.Push(h, v) }
func (h *hpIsPossible) pop() int   { return heap.Pop(h).(int) }

func isPossible1(nums []int) bool {
	lens := map[int]*hpIsPossible{}
	for _, v := range nums {
		if lens[v] == nil {
			lens[v] = new(hpIsPossible)
		}
		if h := lens[v-1]; h != nil {
			prevLen := h.pop()
			if h.Len() == 0 {
				delete(lens, v-1)
			}
			lens[v].push(prevLen + 1)
		} else {
			lens[v].push(1)
		}
	}

	for _, h := range lens {
		if h.IntSlice[0] < 3 {
			return false
		}
	}
	return true
}
