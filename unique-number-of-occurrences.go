// https://leetcode.com/problems/unique-number-of-occurrences

func uniqueOccurrences(arr []int) bool {
	check := map[int]bool{}
	count := map[int]int{}
	a := []int{}
	var res [2000]int
	for _, v := range arr {
		count[v]++
		if !check[v] {
			check[v] = true
			a = append(a, v)
		}
	}
	for _, v := range a {
		if res[count[v]] == 1 {
			return false
		}
		res[count[v]]++
	}
	return true
}