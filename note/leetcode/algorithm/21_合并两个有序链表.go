package algorithm

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{}

	var l1Ptr, l2Ptr, tmp = l1, l2, head
	for {
		if l1Ptr == nil {
			tmp.Next = l2Ptr
			break
		}
		if l2Ptr == nil {
			tmp.Next = l1Ptr
			break
		}

		if l1Ptr.Val > l2Ptr.Val {
			tmp.Next = l2Ptr
			l2Ptr = l2Ptr.Next
		} else {
			tmp.Next = l1Ptr
			l1Ptr = l1Ptr.Next
		}

		tmp = tmp.Next
	}

	return head.Next
}
