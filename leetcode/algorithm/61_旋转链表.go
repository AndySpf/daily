package algorithm

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	tail, l := head, 0
	for tmp := head; ; {
		l++
		if tmp.Next == nil {
			tail = tmp
			break
		}
		tmp = tmp.Next
	}
	if k = k % l; k == 0 {
		return head
	}

	pos := head
	for i := 0; i < l-k-1; i++ {
		pos = pos.Next
	}

	head1 := pos.Next
	pos.Next = nil
	tail.Next = head
	return head1
}
