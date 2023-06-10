package algorithm

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNodes struct {
	nodes []*ListNode
	k     int
}

// 1->2->3->4->5  k=2:2->1->4->3->5 k=3:3->2->1->4->5
// 数组实现方式
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil {
		return head
	}
	nodes := ListNodes{
		nodes: []*ListNode{},
		k:     k,
	}
	next := head
	for {
		tmp := next.Next
		nodes.push(next)
		next = tmp
		if next == nil {
			break
		}
	}

	for i := 0; i < len(nodes.nodes); i++ {
		if i == len(nodes.nodes)-1 {
			break
		}
		nodes.nodes[i].Next = nodes.nodes[i+1]
	}
	return nodes.nodes[0]
}

func (p *ListNodes) push(node *ListNode) {
	node.Next = nil // 清除原有关系

	p.nodes = append(p.nodes, node)
	count := len(p.nodes) / p.k
	if len(p.nodes)%p.k == 0 { // 整数倍触发翻转
		start := (count - 1) * p.k
		end := count * p.k

		tmp := p.nodes[start:end]
		for i := len(tmp)/2 - 1; i >= 0; i-- {
			opp := len(tmp) - 1 - i
			tmp[i], tmp[opp] = tmp[opp], tmp[i]
		}
	}
}

// 1->2->3->4->5  k=2:2->1->4->3->5 k=3:3->2->1->4->5
// 直接翻转
func reverseKGroup1(head *ListNode, k int) *ListNode {
	if k == 1 || head == nil {
		return head
	}

	// 1->2->3
	res := &ListNode{
		Val:  0,
		Next: head,
	}

	curr := head
	p := head  // 记录待翻转的部分的首位，翻转完后会成为下一个待翻转区域的前驱
	pre := res // 记录翻转前的待翻转区域前驱
	for index := 1; curr != nil; index++ {
		if index%k == 0 {
			//触发翻转
			next := curr.Next
			curr.Next = nil // 切断待翻转部分与后继的联系

			p = pre.Next
			pre.Next = nil // 切断待翻转部分与前驱的联系

			pre.Next = reverse(p) // 重建待翻转部分与前驱联系

			p.Next = next // 重建待翻转部分与后继的联系

			pre = p // 生成下一个翻转部分的新的前驱

			curr = p.Next // 更新当前指针位置
			continue
		}
		curr = curr.Next
	}

	return res.Next
}

// nil->1->2->3->nil    nil<-1<-2<-3<-nil
func reverse(curr *ListNode) *ListNode {
	var pre *ListNode // 初始化head的前驱为nil
	for curr != nil {
		next := curr.Next // 记录翻转前的后继
		curr.Next = pre   // 当前的后继变为当前的前驱
		pre = curr        // 前驱变成当前项
		curr = next       // 当前项变成之前记录的翻转前后继
	}
	return pre
}
