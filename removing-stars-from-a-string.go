func removeStars(s string) string {
  var count int
  stack := make([]byte, 0)
  for i := len(s) - 1; i>=0; i-- {
    if s[i] == '*' {
      count++
    } else {
      if count > 0 {
        count-- 
      } else {
        stack = append(stack, s[i])
      }
    }
  }
  res := make([]byte, 0)
  for len(stack) != 0 {
    res = append(res, stack[len(stack) - 1])
    stack = stack[:len(stack) - 1]
  }
  return string(res)
}