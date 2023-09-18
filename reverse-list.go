func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head
	third := head.Next
	for {
		if third.Next == nil {
			third.Next = head
			tmp.Next = nil
			return third
		}
		sec := third
		third = sec.Next
		sec.Next = head
		head = sec
	}
}