// https://leetcode.com/problems/roman-to-integer

class Solution {
public:
    int romanToInt(string s) {
        int res=0;
        char roman[7]={'I','V','X','L','C','D','M'};
        int id[s.length()];
        int value[7]={1,5,10,50,100,500,1000};
        for (int i=0;i<s.length();i++)
        {
            for(int j=0;j<7;j++)
            {
                if (s[i]==roman[j])
                {
                    id[i]=j;
                }
            }
        }
        for(int i=0;i<s.length()-1;i++)
        {
            if(id[i]<id[i+1])
            {
                res-=value[id[i]];
            }
            else
            {
                res+=value[id[i]];
            }
        }
        res+=value[id[s.length()-1]];
        return res;
    }
};