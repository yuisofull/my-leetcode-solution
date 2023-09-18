// https://leetcode.com/problems/two-sum

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> result={-1,-1};
        vector<pair<int,int>> nums2;
        for(int i=0;i<nums.size();i++)
        {
            nums2.push_back(make_pair(nums[i],i));
        }
        sort(nums2.begin(),nums2.end());
        int x=0,y=nums.size()-1;
        while(true)
        {
            if ((nums2[x].first+nums2[y].first)==target)
            {
                result[0]=nums2[x].second;
                result[1]=nums2[y].second;
                break;
            }else{
                if((nums2[x].first+nums2[y].first)<target)
                {
                    x++;
                }else{
                    y--;
                }
            }
            if(x==y){
                break;
            }
        }
        return result;
    }
    
};