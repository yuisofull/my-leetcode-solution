// https://leetcode.com/problems/longest-common-prefix

class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        string s="";
        string res;
        int m=0;
        sort(strs.begin(),strs.end());
        for(int i=0;i<strs.size()-1;i++)
        {
            string test="";
            int h=0;
            for (int j=0;j<strs[i].length();j++)
            {
                
                if (strs[i][j]==strs[i+1][j])
                {
                    h++;
                    string a="";
                    a=strs[i][j];
                    test.insert(test.length(),a);
                }else
                {
                    if(h>m)
                    {
                        m=h;
                        res=test;
                    }
                    break;
                }
            }
        }
        return res;
    }
};