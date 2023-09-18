// https://leetcode.com/problems/increasing-triplet-subsequence

func increasingTriplet(nums []int) bool {
	var a, b, j int
	for i, v := range nums {
		if i == 0 {
			a = v
			b = v
			j = 0
			continue
		}
		if v > a {
			if v < b || j == 0 {
				b = v
				j = i
			}
		}
		if v < a {
			a = v
		}
		if v > b && j != 0 {
			return true
		}
	}
	return false
}