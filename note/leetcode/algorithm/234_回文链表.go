package algorithm

func isPalindromeList(head *ListNode) bool {
	s := []byte{}

	for head != nil {
		s = append(s, uint8(head.Val)+'0')
		head = head.Next
	}

	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}

	return true
}
