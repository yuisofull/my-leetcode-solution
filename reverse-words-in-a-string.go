
func reverseWords(s string) string {
    //s=str.Trim(s," ")
    tmp:=strings.Fields(s)
    res:=""
    for i,_:=range tmp{
        res+=tmp[len(tmp)-i-1]+" "
    }
    return res[:len(res)-1]
}