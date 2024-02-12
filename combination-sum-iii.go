func combinationSum3(k int, n int) [][]int {
    var backTrace func(count, current, sum int)
    var result [][]int
    var res []int
    backTrace = func(count, current, sum int) {
        if count == 0 {
            if sum == n {
                temp := make([]int, k)
                copy(temp, res)
                result = append(result, temp)
            }
            return
        }
        if sum > n {
            return
        }
        for i := current; i <= 9; i++ {
            res = append(res, i)
            backTrace(count - 1, i + 1, sum + i)
            res = res[:len(res)-1]
        } 
    }
    backTrace(k, 1, 0)
    return result
}