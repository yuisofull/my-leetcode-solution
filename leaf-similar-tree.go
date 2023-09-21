/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func re(head *TreeNode, arr []int) []int {
	if head == nil {
		return arr
	}
	if head.Left == nil && head.Right == nil {
		return append(arr, head.Val)
	}
	arr = re(head.Left, arr)
	arr = re(head.Right, arr)
	return arr
}
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var a, b []int
	a = re(root1, a)
	b = re(root2, b)
	return reflect.DeepEqual(a, b)
}