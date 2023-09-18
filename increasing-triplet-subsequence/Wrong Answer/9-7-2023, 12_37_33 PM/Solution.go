// https://leetcode.com/problems/increasing-triplet-subsequence

func increasingTriplet(nums []int) bool {
    for i:=2;i<len(nums);i++{
        if nums[i]>nums[i-1]&&nums[i-1]>nums[i-2]{
            return true
        }
    }
    return false
}