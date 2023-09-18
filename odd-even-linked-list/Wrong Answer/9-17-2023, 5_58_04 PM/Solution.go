// https://leetcode.com/problems/odd-even-linked-list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    //head1 := head
    head2 := head.Next
    i := head
    j := head.Next
    var count int
    for {
        count++
        tmp := i.Next
        i.Next = j.Next
        if j.Next == nil {
            if count % 2 == 1{
                j.Next = head2
            } else {
                i.Next = head2
            }
            break
        }
        i = tmp
        j = j.Next
    }
    return head
}