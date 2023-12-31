// https://leetcode.com/problems/is-subsequence

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	k := 0
	for i := 0; i < len(t); i++ {
		if s[k] == t[i] {
			k++
		}
		if k == len(s) {
			return true
		}
	}
	return false
}