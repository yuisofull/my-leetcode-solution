func kidsWithCandies(candies []int, extraCandies int) []bool {
    var res []bool
    max:=0
    for _,v:=range candies{
        if max<v{
            max=v
        }
    }
    for _,v:=range candies{
        res=append(res,v+extraCandies>=max)
    }
    return res
}