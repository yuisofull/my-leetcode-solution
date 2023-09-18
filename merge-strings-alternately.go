func min(a int, b int) int {
    if a<b {
        return a
    }
    return b
}

func mergeAlternately(word1 string, word2 string) string {
    var res string
    if a:=min(len(word1),len(word2)); a == len(word1){
        for i,_:= range word1 {
            res+=string(word1[i])+string(word2[i])
        }
        for i:=a;i<len(word2);i++ {
            res+=string(word2[i])
        }
        return res
    }else{
        for i,_:= range word2 {
            res+=string(word1[i])+string(word2[i])
        }
        for i:=a;i<len(word1);i++ {
            res+=string(word1[i])
        }
        return res
    }
}