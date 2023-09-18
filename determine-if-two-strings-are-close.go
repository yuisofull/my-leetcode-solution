// https://leetcode.com/problems/determine-if-two-strings-are-close

func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }
    freq1 := make(map[string]int)
    freq2 := make(map[string]int)

    check1 := make(map[string]bool)
    check2 := make(map[string]bool)
    for i:=0; i<len(word1); i++ {
        freq1[string(word1[i])]++
        freq2[string(word2[i])]++
        check1[string(word1[i])] = true
        check2[string(word2[i])] = true
    }
    if !reflect.DeepEqual(check1, check2) {
        return false
    }
    revFreq1 := make(map[int]int)
    revFreq2 := make(map[int]int)
    for _,v := range freq1 {
        revFreq1[v]++
    } 
    for _,v := range freq2 {
        revFreq2[v]++
    }
    return reflect.DeepEqual(revFreq1, revFreq2)