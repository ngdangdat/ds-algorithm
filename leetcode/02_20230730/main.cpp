#include<vector>
#include<iostream>
#include<algorithm>

using namespace std;

class Solution {
public:
    int majorityElement(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        
        if (nums.size() <= 2) return nums[0];
        int l = 0;
        for (int i = 1; i < nums.size(); i++) {
            if (nums[i] == nums[l]) {
                if ((i - l + 1) >= (nums.size() / 2 + nums.size() % 2)) return nums[l];
            }
            if (nums[i] != nums[i - 1]) {
                l = i;
            }
        }
        return nums[l];
    }
};

int main() {
    Solution s = Solution();
    vector<int> inputs = {6, 6, 6, 7, 7};
    // vector<int> inputs = {3, 2, 3};
    int result = s.majorityElement(inputs);
    printf("==========Sorted\n");
    for (int i = 0; i < inputs.size(); i++) {
        printf("nums[%d]=%d\n", i, inputs[i]);
    }
    printf("==========Result: %d\n", result);
}
