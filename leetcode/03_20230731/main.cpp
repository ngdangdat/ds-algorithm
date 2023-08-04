#include<iostream>
#include<vector>


using namespace std;

class Solution {
public:
    int jump(vector<int>& nums) {
        if (nums.size() < 2) return 0;
        if (nums[0] == 0) return -1; // never happens

        int jump_count = 0;
        for (int i = 0; i < nums.size(); i++) {
            if ((i + nums[i]) >= (nums.size() - 1)) {
                jump_count += 1;
                break;
            }
            int max_index = i + 1;
            for (int j = i + 1; j <= i + nums[i]; j++) {
                if (nums[j] + j >= nums[max_index] + max_index) {
                    max_index = j;
                }
            }
            jump_count += 1;
            if (max_index > i) i = max_index - 1;
        }
        return jump_count;
    }
};

int main() {
    /**
     * @brief Input: nums = [2,3,1,1,4]
              Output: 2
              Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
     * 
     */
    vector<int> inputs = {10,9,8,7,6,5,4,3,2,1,1,0};
    int expected = 2;

    Solution sol = Solution();
    printf("result=[{%d}], expected=[{%d}]\n", sol.jump(inputs), expected);
}
