// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element

func longestSubarray(nums []int) int {
    zeroes := 0
    max := 0
    for l,r := 0,0; r < len(nums); r++ {
        if nums[r] == 0 {
            zeroes++
        }

        if zeroes > 1 {
            if nums[l] == 0 {
                zeroes--
            }

            l++
        }

        if max < r-l+1 {
            max = r-l+1
        }
    }

    return max - 1
}