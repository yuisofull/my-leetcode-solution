func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxDepth(root.Left) + 1
	r := maxDepth(root.Right) + 1
	return func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}(l, r)
}