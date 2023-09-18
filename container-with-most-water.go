// https://leetcode.com/problems/container-with-most-water

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	var m int
	if height[l] < height[r] {
		m = height[l] * (r - l)
	} else {
		m = height[r] * (r - l)
	}

	for l < r {
		if height[l] < height[r] {
			m = max(m, height[l]*(r-l))
			l++
		} else {
			m = max(m, height[r]*(r-l))
			r--
		}
	}
	return m
}