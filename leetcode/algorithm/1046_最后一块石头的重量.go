package algorithm

import (
	"container/heap"
	"sort"
)

func lastStoneWeight(stones []int) int {
	for len(stones) >= 2 {
		sort.Ints(stones)
		l := len(stones)
		if stones[l-1] == stones[l-2] {
			stones = stones[:l-2]
		}
		if stones[l-1] < stones[l-2] {
			stones = stones[:l-1]
			stones[l-2] = stones[l-1] - stones[l-2]
		}
	}
	if len(stones) == 0 {
		return 0
	}
	return stones[0]
}

type hpstore struct{ sort.IntSlice }

func (h hpstore) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hpstore) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hpstore) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hpstore) push(v int) { heap.Push(h, v) }
func (h *hpstore) pop() int   { return heap.Pop(h).(int) }

func lastStoneWeight1(stones []int) int {
	q := &hpstore{stones}
	heap.Init(q)
	for q.Len() > 1 {
		x, y := q.pop(), q.pop()
		if x > y {
			q.push(x - y)
		}
	}
	if q.Len() > 0 {
		return q.IntSlice[0]
	}
	return 0
}
