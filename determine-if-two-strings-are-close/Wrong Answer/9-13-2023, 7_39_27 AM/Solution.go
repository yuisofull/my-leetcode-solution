// https://leetcode.com/problems/determine-if-two-strings-are-close

import "sort"

// Check if sorted frequency of characters are the same.
// Check if set of characters are the same
func closeStrings(word1 string, word2 string) bool {
    charSet1, charCount1 := getCharSetAndSortedCharCount(word1)
    charSet2, charCount2 := getCharSetAndSortedCharCount(word2)

    if charSet1 != charSet2 {
        return false
    }

    for i := 0; i < 26; i++ {
        if charCount1[i] != charCount2[i] {
            return false
        }
    }

    return true
}

func getCharSetAndSortedCharCount(word string) (int, []int) {
    charSet := 0
    sortedCount := make([]int, 26)
    for _, c := range word {
        cidx := int(c - 'a')

        charSet |=  cidx
        sortedCount[cidx] += 1
    }

    sort.Slice(sortedCount, func(i, j int) bool {
        return sortedCount[i] < sortedCount[j]
    })

    return charSet, sortedCount
}