// https://leetcode.com/problems/delete-the-middle-node-of-a-linked-list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    var n int
    for i := head; i != nil; i = i.Next {
        n++ 
    }
    p := head
    var l *ListNode
    var r *ListNode
    for i := 1; i <= (n/2 + 1); i++ {
        if i == n/2 {
            l = p
        }
        if i == (n/2 + 1) {
            r = p.Next
        }
        p = p.Next
    }
    l.Next = r
    return head
}