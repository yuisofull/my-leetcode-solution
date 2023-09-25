/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxLevelSum(root *TreeNode) int {
	res, level, max := 1, 1, root.Val
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		n := len(queue)
		var sum int
		added := false
		for i := 0; i < n; i++ {
			if queue[0] == nil {
				//queue = queue[1:]
				break
			}
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
				sum += queue[0].Left.Val
				added = true
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
				sum += queue[0].Right.Val
				added = true
			}
			//sum -= queue[0].Val
			queue = queue[1:]
		}
		level++
		if max < sum && added {
			max = sum
			res = level
		}
	}
	return res
}