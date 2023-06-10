package algorithm

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	var fast, slow *ListNode
	fast = head.Next
	slow = head
	for {
		if fast == slow {
			return true
		}
		if fast == nil {
			return false
		}

		if fast.Next == nil {
			return false
		} else {
			fast = fast.Next.Next
		}

		slow = slow.Next
	}
}
