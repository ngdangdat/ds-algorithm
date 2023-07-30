#include<vector>
#include <iostream>

using namespace std;


class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        if (nums.size() <= 1) return nums.size();
        int count = 0;
        int index = 0;
        for (int i = 0; i < nums.size(); i++) {
            if (i == 0) {
                index += 1;
                count += 1;
                continue;
            }
            if (nums[i] == nums[i - 1] && count < 2) {
                count += 1;
                nums[index] = nums[i];
                index += 1;
            }
            if (nums[i] != nums[i - 1]) {
                count = 1;
                nums[index] = nums[i];
                index += 1;
            }
        }
        return index;
    }
};

class Solution2 {
    /**
     * @brief Improved one
     * 
     */
public:
    int removeDuplicates(vector<int>& nums) {
        if (nums.size() <= 2) return nums.size();
        int r = 1;
        // will be wrong if we start from 0
        for (int i = 2; i < nums.size(); i++) {
            if (nums[r] == nums[i] && nums[r - 1] == nums[i]) continue;
            nums[++r] = nums[i];
        }
        return r + 1;
    }
};

int main() {
    Solution2 sol = Solution2();
    vector<int> inputs = {1, 2, 2};
    int index = sol.removeDuplicates(inputs);
    for (int i = 0; i < index; i++) {
        printf("nums[%d]=%d\n", i, inputs[i]);
    }
}
