/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	var res int
	var dfs func(*TreeNode, int, bool)
	var max func(int, int) int
	max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	dfs = func(head *TreeNode, count int, left bool) {
		if head == nil {
			res = max(count, res)
			return
		}
		if left {
			dfs(head.Right, count+1, false)
			dfs(head.Left, 0, true)
		} else {
			dfs(head.Left, count+1, true)
			dfs(head.Right, 0, false)
		}
		return
	}
	dfs(root.Left, 0, true)
	dfs(root.Right, 0, false)
	return res
}