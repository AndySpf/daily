package algorithm

// 输入: 1->2->3->4->5->6->7->NULL   1,3,2,4,5
// 输出: 1->3->5->7->2->4->6->NULL
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	odd, even := head, head.Next // odd指向当前排列好的奇数队列最后一个，even指向当前已排列好的偶数队列最后一个
	for even != nil && odd != nil && even.Next != nil {
		nextOdd := even.Next
		even.Next = even.Next.Next

		nextOdd.Next = odd.Next
		odd.Next = nextOdd

		odd = odd.Next
		even = even.Next
	}
	return head
}
