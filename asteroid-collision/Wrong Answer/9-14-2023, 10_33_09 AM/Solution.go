// https://leetcode.com/problems/asteroid-collision

func asteroid(stack []int, v int) ([]int, int) {
    if len(stack) == 0 {
      stack = append(stack, v)
      return stack, v
    }
    if v < 0 {
      if stack[len(stack) - 1] > 0 {
        i := stack[len(stack) - 1] + v
        switch {
            case i == 0:
                stack = stack [:len(stack) - 1]
                return stack, v
            case i < 0:
                stack, _ = asteroid(stack[:len(stack) - 1], v)
                if (stack[len(stack) - 1] + v) < 0 {
                    return append(stack, v), v
                }
                return stack, v 
                
            default:
                return stack, v
        }
      } else {
        stack = append(stack, v)
        return stack, v
      }
    } else {
        if stack[len(stack) - 1] < 0 {
            i := stack[len(stack) - 1] + v
            switch {
                case i == 0:
                    stack = stack [:len(stack) - 1]
                    return stack, v
                case i > 0:
                    stack, _ = asteroid(stack[:len(stack) - 1], v)
                    if (stack[len(stack) - 1] + v) > 0 {
                        return append(stack, v), v
                    }
                    return stack, v
            }
        } else {
            stack = append(stack, v)
            return stack, v
      }
    }
    return stack, v
}

func asteroidCollision(asteroids []int) []int {
  stack := make([]int, 0)
  for _, v := range asteroids {
    stack, _ = asteroid(stack, v)
  }
  return stack
}