package algorithm

func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	sortedHead, unsortedHead := head, head.Next
	sortedHead.Next = nil
	for unsortedHead != nil {
		p, o := sortedHead, unsortedHead
		unsortedHead = unsortedHead.Next
		o.Next = nil
		for {
			if o.Val < p.Val { // 比第一个小
				o.Next = p
				sortedHead = o
				break
			} else if p.Next == nil {
				p.Next = o
				break
			} else if o.Val >= p.Val && o.Val < p.Next.Val {
				nextItem := p.Next
				p.Next = o
				o.Next = nextItem
				break
			}

			p = p.Next
		}
	}
	return sortedHead
}
