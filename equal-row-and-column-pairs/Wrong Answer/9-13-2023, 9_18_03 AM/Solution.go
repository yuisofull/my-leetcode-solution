// https://leetcode.com/problems/equal-row-and-column-pairs

func equalPairs(grid [][]int) int {
  m := map[string]int{}
  var res int
  for _, row := range grid {
    var buffer bytes.Buffer
    for _, val := range row {
      num := strconv.Itoa(val)
      buffer.WriteString(num)
    }
    m[buffer.String()]++ 
  }
  for i, _ := range grid {
    var buffer bytes.Buffer
    for j, _ := range grid {
      num := strconv.Itoa(grid[j][i])
      buffer.WriteString(num)
    }
    if m[buffer.String()] >= 1 {
      res += m[buffer.String()]
    }
  }
  return res
}