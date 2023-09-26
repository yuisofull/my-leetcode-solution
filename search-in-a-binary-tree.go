/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root.Val == val {
		return root
	}
	var queue = []*TreeNode{root}
	for len(queue) > 0 {
		var children []*TreeNode
		for _, node := range queue {
			if node.Left != nil {
				if node.Left.Val == val {
					return node.Left
				}
				children = append(children, node.Left)
			}
			if node.Right != nil {
				if node.Right.Val == val {
					return node.Right
				}
				children = append(children, node.Right)
			}
		}
		queue = children
	}
	return nil
}
