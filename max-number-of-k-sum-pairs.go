// https://leetcode.com/problems/max-number-of-k-sum-pairs

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	l := 0
	r := len(nums) - 1
	res := 0
	for l < r {
		if (nums[l] + nums[r]) == k {
			res++
			l++
			r--
		}
		if (nums[l] + nums[r]) > k {
			r--
		}
		if (nums[l] + nums[r]) < k {
			l++
		}
	}
	return res
}