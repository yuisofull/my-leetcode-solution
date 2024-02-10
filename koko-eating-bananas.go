func minEatingSpeed(piles []int, h int) int {
    l := 1
    for r := slices.Max(piles); l <= r; {
        mid := (l + r) / 2
        var sum int
        for _, pile := range piles {
            sum += int(math.Ceil(float64(pile) / float64(mid)))
            if sum > h {
                break
            }
        }
        if sum > h {
            l = mid + 1
        }else{
            r = mid - 1
        }
    }
    return l
}