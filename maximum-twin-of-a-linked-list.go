func pairSum(head *ListNode) int {
	max := 0
	var stack []int
	for head != nil {
		if head.Next == nil {
			if max < (stack[0] + head.Val) {
				max = stack[0] + head.Val
				stack = stack[1:]
			}
			for len(stack) != 0 {
				if max < (stack[0] + stack[len(stack)-1]) {
					max = stack[0] + stack[len(stack)-1]
				}
				stack = stack[1 : len(stack)-1]
			}
		}
		stack = append(stack, head.Val)
		head = head.Next
	}
	return max
}