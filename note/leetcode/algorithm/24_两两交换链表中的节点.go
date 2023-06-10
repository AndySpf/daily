package algorithm

// before  after   nextBefore nextAfter
// 1,      2,      3,         4
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var before, after = head, head.Next
	if head.Next == nil {
		return head
	} else {
		head = head.Next
	}
	for {
		var nextBefore, nextAfter *ListNode
		if after.Next != nil {
			nextBefore = after.Next
			if nextBefore.Next != nil {
				nextAfter = after.Next.Next
			}
		}

		if nextAfter == nil {
			before.Next = nextBefore
		} else {
			before.Next = nextAfter
		}
		after.Next = before

		before = nextBefore
		after = nextAfter

		if before == nil || after == nil {
			break
		}
	}
	return head
}
