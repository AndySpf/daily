package algorithm

// 如果快指针追上了慢指针，则快指针走的距离一定是慢指针的2倍(因为距离有倍数关系，快慢指针都从head出发，区别题目141)
// a = 没进入环的距离 b = 从入环点到相遇点距离 c = 从相遇点到下一次经过入环点的距离
// 相遇时，快指针走过的距离 = nd + b + a
// 慢指针走过的距离    = a + b
// 2(a + b) = a + nd + b   slow入环后第一圈必定相遇。因为fast是slow移速的2倍，slow入环开始，走满一圈，fast已经走满2圈了，肯定相遇
// a = n(b+c) - b = (n-1)(b+c) + c
// 即如果有环存在，则一个指针从起始点开始，一个指针从相遇点开始继续走。必能相遇，且相遇时走过的距离就是入环点
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var fast, slow *ListNode = head, head
	isCircle := false

	for {
		if fast == nil || fast.Next == nil {
			break
		}

		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			isCircle = true
			break
		}
	}

	if !isCircle {
		return nil
	}

	tmp := head
	for {
		if fast == tmp {
			return fast
		}
		fast = fast.Next
		tmp = tmp.Next
	}
}
