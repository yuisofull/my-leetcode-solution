// https://leetcode.com/problems/find-the-highest-altitude

func largestAltitude(gain []int) int {
    var sum, max int
    for _, v := range gain {
        sum += v
        if sum > max {
            max = sum
        }
    }

    return max
}