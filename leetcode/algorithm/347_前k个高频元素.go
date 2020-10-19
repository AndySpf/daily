package algorithm

import (
	"container/heap"
)

func topKFrequent(nums []int, k int) []int {
	counter := map[int]int{}
	arr := make([][]int, len(nums)+1)
	res := make([]int, 0, k)
	for i := range nums {
		if _, ok := counter[nums[i]]; ok {
			counter[nums[i]]++
		} else {
			counter[nums[i]] = 1
		}
	}

	for key, v := range counter {
		arr[v] = append(arr[v], key)
	}

	for i := len(nums); i >= 0; i-- {
		if len(arr[i]) != 0 {
			for _, item := range arr[i] {
				res = append(res, item)
				if len(res) == cap(res) {
					return res
				}
			}
		}
	}
	return res
}

func topKFrequent1(nums []int, k int) []int {
	counter := map[int]int{}
	for i := range nums {
		if _, ok := counter[nums[i]]; ok {
			counter[nums[i]]++
		} else {
			counter[nums[i]] = 1
		}
	}

	h := new(myHeap)
	heap.Init(h)

	for key, v := range counter {
		heap.Push(h, heapItem{
			count: v,
			value: key,
		})

		if h.Len() > k {
			heap.Pop(h) // pop出最小的
		}
	}
	res := make([]int, 0, k)
	for {
		item := heap.Pop(h).(heapItem)
		res = append(res, item.value)
		if h.Len() == 0 {
			break
		}
	}
	return res
}

//     2             -> 1 -> 丢弃
//  3     4          -> 3 -> 丢弃2
type myHeap []heapItem

type heapItem struct {
	count int
	value int
}

func (p myHeap) Len() int {
	return len(p)
}

func (p myHeap) Less(i, j int) bool {
	return p[i].count < p[j].count
}

func (p myHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *myHeap) Pop() interface{} {
	old := *p
	length := len(*p)
	x := old[length-1]
	old = old[:length-1]
	*p = old
	return x
}

func (p *myHeap) Push(x interface{}) {
	*p = append(*p, x.(heapItem))
}
