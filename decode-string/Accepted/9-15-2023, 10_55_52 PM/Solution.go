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
            var quantityStr string
			for len(stack) > 0 && (stack[len(stack) - 1] >= '0') && (stack[len(stack) - 1] <= '9') {
				quantityStr = string(stack[len(stack)-1]) + quantityStr
				stack = stack[:len(stack)-1]
			}
			quantity, _ := strconv.Atoi(quantityStr)
            for i := 0; i < quantity ; i++ {
                stack = append(stack, tmpString...)
            }
        } else {
            stack = append(stack, s[i])
        }
    }
    return string(stack)
}