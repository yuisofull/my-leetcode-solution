// https://leetcode.com/problems/reverse-words-in-a-string


func reverseWords(s string) string {
    //str:=[]byte(s)
    r:=len(s)
    res:=""
    for i:=len(s)-1;i>=0;i--{
        if string(s[i])==" "{
            l:=i+1
            if l!=r{
                res+=s[l:r]+" "
            }
            r=i
        }
    }
    if string(s[0])!=" "{
        res+=s[0:r+1]
    }
    return res[:len(res)-1]
}