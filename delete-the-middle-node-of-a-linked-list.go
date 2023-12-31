// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list

func deleteMiddle(head *ListNode) *ListNode {
	var n int
	for i := head; i != nil; i = i.Next {
		n++
	}
	p := head
	var l, r *ListNode
	for i := 1; i <= (n/2 + 1); i++ {
		if i == n/2 {
			l = p
		}
		if i == (n/2 + 1) {
			r = p.Next
		}
		p = p.Next
	}
	if l != nil {
		l.Next = r
	} else {
		return nil
	}
	return head
}