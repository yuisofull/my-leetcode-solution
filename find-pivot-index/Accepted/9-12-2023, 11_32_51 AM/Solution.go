// https://leetcode.com/problems/find-pivot-index

func pivotIndex(nums []int) int {
    var l, r int
    for _, v := range nums {
        r += v
    }

    for i, v := range nums {
        r -= v
        if r == l {
            return i
        }
        l += nums[i]
    }
    return -1
}