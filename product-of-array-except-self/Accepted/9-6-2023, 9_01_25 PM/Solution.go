// https://leetcode.com/problems/product-of-array-except-self

func productExceptSelf(nums []int) []int {
    product:=1
    zero:=0
    onlyNumber:=0
    for i,v:=range nums{
        if v==0{
            zero++
            onlyNumber=i
            continue
        }
        product*=v
    }
    if zero == 1{
        res:=make([]int, len(nums))
        res[onlyNumber]=product
        return res
    }
    if zero>1{
        return make([]int, len(nums))
    }
    var res []int
    for _,v:=range nums{
        res=append(res, product/v)      
    }
    return res
}