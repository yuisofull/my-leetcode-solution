/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		n := len(queue)
		if queue[n-1] != nil {
			res = append(res, queue[n-1].Val)
		}
		for i := 0; i < n; i++ {
			if queue[0] == nil {
				queue = queue[1:]
				break
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			queue = queue[1:]
		}
	}
	return res
}