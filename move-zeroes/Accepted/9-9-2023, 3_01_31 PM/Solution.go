// https://leetcode.com/problems/move-zeroes

func moveZeroes(nums []int)  {
    nextIndex := 0

    for i, num := range nums {
        if num != 0 {
            nums[i], nums[nextIndex] = nums[nextIndex], nums[i]
            nextIndex++
        }
    }
}