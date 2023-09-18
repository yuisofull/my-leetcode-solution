func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    head2 := head.Next
    i := head
    j := head.Next
    var count int
    for {
        count++
        tail := i
        tmp := i.Next
        i.Next = j.Next
        if j.Next == nil {
            if count % 2 == 1 {
                tail.Next = head2
            } else {
                i = tmp
                i.Next = head2
            }
            break
        }
        i = tmp
        j = j.Next
    }
    return head
}