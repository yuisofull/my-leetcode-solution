// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxVowels(s string, k int) int {
	sum := 0
	dict := map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1}
	for i := 0; i < k; i++ {
		sum += dict[string(s[i])]
	}
	res := sum
	for i := 1; i < len(s)-k+1; i++ {
		sum = sum - dict[string(s[i-1])] + dict[string(s[i+k-1])]
		res = max(res, sum)
		if res == k {
			return res
		}
	}
	return res
}