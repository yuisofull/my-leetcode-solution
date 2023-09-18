// https://leetcode.com/problems/max-consecutive-ones-iii

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func longestOnes(nums []int, k int) int {
    sum := 0
    res := 0 
    pos := 0
    count := k
    for i, v := range nums {
        if v == 1 {
            sum++
        } else {
            if count > 0 {
                count--
                sum++
            } else {
                res = max(res, sum)
                for nums[pos] != 0 && pos < len(nums){
                    pos++
                }
                sum = i - pos
                pos++
            }  
        }
    }
    return res
}