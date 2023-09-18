// https://leetcode.com/problems/move-zeroes

func moveZeroes(nums []int)  {
    k := 0

    for i,v := range nums {
        if(v != 0) {
            if ( i!=k ){
                nums[i], nums[k] = nums[k], nums[i]
            } 
            k++
        }
    }
}