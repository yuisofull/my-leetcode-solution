/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func DFS(head *TreeNode, max int, count *int) {
	if head == nil {
		return
	}
	if head.Val >= max {
		*count++
		max = head.Val
	}
	DFS(head.Left, max, count)
	DFS(head.Right, max, count)
}
func goodNodes(root *TreeNode) int {
	var count int
	max := root.Val
	DFS(root, max, &count)
	return count
}