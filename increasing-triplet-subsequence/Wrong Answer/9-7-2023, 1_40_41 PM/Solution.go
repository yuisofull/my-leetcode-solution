// https://leetcode.com/problems/increasing-triplet-subsequence

func increasingTriplet(nums []int) bool {
    var a,b int
    for i,v:=range nums{
        if i==0{
            a=v
            b=v
            continue
        }
        if v>a{
            if v<b || b==a{
                b=v
            }
        }
        if v<a{
            a=v
        }
        if v>b && b!=a{
            return true
        }
    }
    return false
}