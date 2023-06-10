package algorithm

// #->1->2->3->4->5
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{
		Val:  0,
		Next: head,
	}
	var start, end = dummyNode, dummyNode
	for i := 0; i < n; i++ {
		if end.Next == nil {
			return head
		}
		end = end.Next
	}

	for {
		if end.Next == nil {
			start.Next = start.Next.Next
			break
		}

		start = start.Next
		end = end.Next
	}
	return dummyNode.Next
}
