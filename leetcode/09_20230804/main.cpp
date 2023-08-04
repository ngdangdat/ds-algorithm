/**
 * @file main.cpp
 * @author your name (you@domain.com)
 * @brief https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/?envType=study-plan-v2&envId=top-interview-150
 * @version 0.1
 * @date 2023-08-04
 * 
 * @copyright Copyright (c) 2023
 * 
 */
#include<iostream>
#include<vector>
#include<algorithm>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        vector<int> res(2, 0);
        int max_target = target;
        if (numbers[0] < 0) {
            max_target = target - numbers[0];
        }
        for (int l = 0, r = numbers.size() - 1; l < r;) {
            while(numbers[r] > max_target) {
                r -= 1;
                continue;
            }
            if (numbers[l] + numbers[r] == target) {
                res[0] = l + 1;
                res[1] = r + 1;
                break;
            } else if (numbers[l] + numbers[r] > target) {
                r -= 1;
            } else {
                l += 1;
            }
        }
        return res;
    }
};

int main() {
    vector<int> inputs1 = {2,3,4};
    int target1 = 6;
    Solution s = Solution();
    vector<int> res1 = s.twoSum(inputs1, target1);
    printf("1 %d %d\n", res1[0], res1[1]);

}
