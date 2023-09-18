// https://leetcode.com/problems/find-pivot-index

func pivotIndex(nums []int) int {
    var l, r int
    for _, v := range nums {
        r += v
    }

    for i, v := range nums {
        if i == 0 {
            r -= v
            if r == l {
                return i
            }
            continue
        }
        r -= v
        l += nums[i - 1]
        if r == l {
            return i
        }
    }
    return -1
}