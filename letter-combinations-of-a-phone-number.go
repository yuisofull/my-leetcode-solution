func letterCombinations(digits string) []string {
    letters := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
    var result []string
    var dfs func(string, string)
    if digits == "" {
        return []string{}
    }
    dfs = func(current, res string){
        if len(current) == 0 {
            result = append(result, res)
            return
        }
        for _, letter := range letters[current[0] - 50] {
            dfs(current[1:], fmt.Sprintf("%s%c", res, letter))
        }
    }
    dfs(digits, "")
    return result
}