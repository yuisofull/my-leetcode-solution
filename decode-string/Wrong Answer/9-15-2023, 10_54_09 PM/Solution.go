// https://leetcode.com/problems/decode-string

func decodeString(s string) string {
    stack := make([]byte, 0)
    for i := 0; i < len(s); i++ {
        if s[i] == ']' {
            count := len(stack) - 1
            for stack[count] != '[' {
                count --                
            }
            tmpString := make([]byte, len(stack[count + 1:]))
            copy(tmpString, stack[count + 1:]) 
			stack = stack[:count]
            quantity := 0
            for len(stack) >0 && (stack[len(stack) - 1] >= '0') && (stack[len(stack) - 1] <= '9') {
                quantity = quantity * 10 + int(stack[len(stack) - 1] - 48)
				stack = stack [:len(stack) - 1]
            }
            for i := 0; i < quantity ; i++ {
                stack = append(stack, tmpString...)
            }
        } else {
            stack = append(stack, s[i])
        }
    }
    return string(stack)
}