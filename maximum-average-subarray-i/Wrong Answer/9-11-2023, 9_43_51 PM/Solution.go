// https://leetcode.com/problems/maximum-average-subarray-i

func m(i, j int) int {
  if i > j {
    return i
  }
  return j
}
func findMaxAverage(nums []int, k int) float64 {
  var max int
  for i := 0; i < k; i++ {
    max += nums[i]
  }
  for i := 1; i < len(nums) - k + 1; i++ {
    tmp := max - nums[i-1] + nums[i+k-1]
    max = m(tmp, max)
  }
  return float64(max)/float64(k)
}