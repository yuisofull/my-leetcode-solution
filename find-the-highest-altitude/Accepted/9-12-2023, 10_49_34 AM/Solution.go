// https://leetcode.com/problems/find-the-highest-altitude

func largestAltitude(gain []int) int {
    max, sum := 0, 0
    for _, v := range gain {
        sum += v
        if sum > max {
            max = sum
        }
    }
    
    return max
}