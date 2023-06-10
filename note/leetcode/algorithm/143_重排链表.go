package algorithm

import "fmt"

// 找中点，翻转后半部分，合并链表 1->2->3->4->5 => 1->2->3 5->4 => 1->5->2->4->3.
// 或者用线性表存一次
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	slow, fast := head, head
	for {
		if fast.Next == nil { // 奇数
			break
		} else if fast.Next.Next == nil { // 偶数
			break
		} else {
			fast = fast.Next.Next
		}
		slow = slow.Next
	}
	reverse := reverseListNode(slow.Next)
	slow.Next = nil
	mergeListNode(head, reverse)
	readNode(head)
}

func readNode(node *ListNode) {
	for {
		if node == nil {
			break
		}
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Print("\n")
}

func reverseListNode(node *ListNode) *ListNode {
	cur := node
	next := node.Next
	cur.Next = nil
	for {
		tmp := next
		if tmp == nil {
			break
		}
		next = tmp.Next
		tmp.Next = cur
		cur = tmp
	}
	return cur
}

// 1->2->3
// 5->4
func mergeListNode(list1, list2 *ListNode) { // len(list2) >= len(list1)
	var list1Ptr, list2Ptr = list1, list2
	for list2Ptr != nil {
		list1Tmp := list1Ptr.Next
		list2Tmp := list2Ptr.Next

		list1Ptr.Next = list2Ptr
		list1Ptr = list1Tmp

		list2Ptr.Next = list1Ptr
		list2Ptr = list2Tmp
	}
}
