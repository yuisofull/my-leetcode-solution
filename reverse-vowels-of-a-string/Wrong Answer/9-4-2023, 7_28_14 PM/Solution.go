// https://leetcode.com/problems/reverse-vowels-of-a-string

func reverseVowels(s string) string {
    arr := []int{}
    res := []byte(s)
    vowels := [...]string{"a","e","i","o","u"}
    for i,v:=range s{
        for _,vow:=range vowels{
            if string(v)==vow{
                arr=append(arr, i)
            }
        }
    }
    for i,v:=range arr{
        res[v]=s[arr[len(arr)-1-i]]
    }
    return string(res)
}