package algorithm

import (
	"container/heap"
	"sort"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并
func mergeKLists1(lists []*ListNode) *ListNode {
	pq := []int{}
	for i := range lists {
		if lists[i] == nil {
			continue
		}
		node := lists[i]
		for {
			pq = append(pq, node.Val)
			if node.Next == nil {
				break
			}
			node = node.Next
		}
	}
	if len(pq) == 0 {
		return nil
	}
	sort.Ints(pq)

	res := &ListNode{}
	tmp := res
	for i := range pq {
		tmp.Val = pq[i]
		if i != len(pq)-1 {
			next := &ListNode{
				Val: pq[i+1],
			}
			tmp.Next = next
			tmp = tmp.Next
		}
	}
	return res
}

// 优先队列
func mergeKLists(lists []*ListNode) *ListNode {
	pq := priorityQueue{}
	for i := range lists {
		if lists[i] == nil {
			continue
		}
		node := lists[i]
		for {
			pq = append(pq, node)
			if node.Next == nil {
				break
			}
			node = node.Next
		}
	}
	if len(pq) == 0 {
		return nil
	}
	heap.Init(&pq)

	res := heap.Pop(&pq).(*ListNode)
	res.Next = nil
	// 从优先队列中逐个读取
	tmp := res
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*ListNode)
		item.Next = nil
		tmp.Next = item
		tmp = item
	}
	return res
}

type priorityQueue []*ListNode

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*p = append(*p, item)
}

func (p *priorityQueue) Pop() interface{} {
	old := *p // 拷贝一份，将拷贝的old处理后赋值给p地址指向的值
	item := old[len(*p)-1]
	*p = old[0 : len(*p)-1]
	return item
}
