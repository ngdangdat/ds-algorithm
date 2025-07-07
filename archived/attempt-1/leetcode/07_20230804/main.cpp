/**
 * @file main.cpp
 * @author your name (you@domain.com)
 * @brief https://leetcode.com/problems/product-of-array-except-self/submissions/?envType=study-plan-v2&envId=top-interview-150
 * @version 0.1
 * @date 2023-08-04
 * 
 * @copyright Copyright (c) 2023
 * 
 */

#include<iostream>
#include<vector>

using namespace std;

class Solution {
public:
    vector<int> productExceptSelf(vector<int>& nums) {
        vector<int> products(nums.size(), 0);
        int product = 1;
        int zero_cnt = 0;
        for (int i = 0; i < nums.size(); i++) {
            if (nums[i] == 0) {
                zero_cnt += 1;
                continue;
            }
            product *= nums[i];
        }
        if (zero_cnt > 1) return products;
        for (int i = 0; i < nums.size(); i++) {
            if (zero_cnt > 0) {
                if (nums[i] == 0) {
                    products[i] = product;
                } else {
                    products[i] = 0; 
                }
            } else {
                products[i] = product / nums[i];
            }
        }
        return products;
    }
};

int main() {
    vector<int> inputs1 = {1,2,3,4};
    vector<int> inputs2 = {-1,1,0,-3,3};
    Solution sol = Solution();
    vector<int> output1 = sol.productExceptSelf(inputs1);
    vector<int> output2 = sol.productExceptSelf(inputs2);

    for (int i = 0; i < output1.size(); i++) {
        printf("%d, ", output1[i]);
    }
    printf("\n");

    for (int i = 0; i < output2.size(); i++) {
        printf("%d, ", output2[i]);
    }
    printf("\n");

    return 0;
}
