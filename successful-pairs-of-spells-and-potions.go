func successfulPairs(spells []int, potions []int, success int64) []int {
    sort.Ints(potions)
    result := make([]int, len(spells))
    for i, spell := range spells {
        for low, hi := 0, len(potions) - 1; ; {
            mid := (low + hi)/2
            if mid == 0 {
                if int64(spell) * int64(potions[mid]) >= success {
                    result[i] += len(potions)
                } else {
                    if int64(spell) * int64(potions[mid + 1]) >= success && len(potions) == 2 {
                        result[i]++
                    }
                }
                break
            }
            if mid == len(potions) - 1 {
                if int64(spell) * int64(potions[mid]) >= success{
                    result[i]++ 
                }
                break
            }
            if int64(spell) * int64(potions[mid]) >= success {
                if int64(spell) * int64(potions[mid - 1]) >= success {
                    hi = mid
                } else {
                    result[i] += len(potions) - mid 
                    break
                }
            } else {
                if int64(spell) * int64(potions[mid + 1]) < success {
                    low = mid + 1
                } else {
                    result[i] += len(potions) - mid - 1
                    break
                }
            }
        }
    }
    return result
}
