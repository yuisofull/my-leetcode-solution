// https://leetcode.com/problems/reverse-words-in-a-string

import str "strings"
func reverseWords(s string) string {
    //s=str.Trim(s," ")
    tmp:=str.Fields(s)
    res:=""
    for i,_:=range tmp{
        res+=str.Trim(tmp[len(tmp)-i-1]," ")+" "
    }
   fmt.Println(tmp)
    return res[:len(res)-1]
}