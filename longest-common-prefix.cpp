// https://leetcode.com/problems/longest-common-prefix

class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {

        string res="";
        int m=0;
        int wc=99999;
        string best;
        for(int i=0;i<strs.size();i++)
        {
            if(strs[i].length()<wc)
            {
                wc=strs[i].length();
                best=strs[i];
            }
        }
        for(int i=0;i<best.length();i++)
        {
            bool ok=true;
            for(int j=0;j<strs.size();j++)
            {

                if(strs[j][i]!=best[i])
                {
                   ok=false;
                }
            }
            if(ok==true)
            {
                m++;
            }else
            {

                break;
            }
        }
        if(m>0)
        {
            res=best.substr(0,m);
        }
        return res;
    }
};