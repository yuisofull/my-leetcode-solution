/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	var find func(**TreeNode)
	find = func(head **TreeNode) {
		if *head == nil {
			return
		}
		if (*head).Val == key {
			switch {
			case (*head).Right == nil && (*head).Left == nil:
				*head = nil
			case (*head).Right == nil:
				*head = (*head).Left
			case (*head).Left == nil:
				*head = (*head).Right
			default:
				tmp := &(*head).Right
				for (*tmp).Left != nil {
					tmp = &(*tmp).Left
				}
				if tmp != &(*head).Right {
					(*head).Val = (*tmp).Val
					*tmp = (*tmp).Right
				} else {
					(*tmp).Left = (*head).Left
					*head = *tmp
				}
			}
			return
		}
		if (*head).Val > key {
			find(&(*head).Left)
		} else {
			find(&(*head).Right)
		}
	}
	find(&root)
	return root
}