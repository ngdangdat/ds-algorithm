/*/**
 * @file main.cpp
 * @author your name (you@domain.com)
 * @brief https://leetcode.com/problems/h-index/submissions/?envType=study-plan-v2&envId=top-interview-150
 * @version 0.1
 * @date 2023-08-03
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
    int hIndex(vector<int>& citations) {
        sort(citations.begin(), citations.end());
        for (int i = 0; i < citations.size(); i++) {
            if (citations[i] >= (citations.size() - i)) {
                return citations.size() - i;
            }
        }
        return 0;
    }
};


int main() {
    vector<int> inputs1 = {3,0,6,1,5};
    int expected1 = 3;
    vector<int> inputs2 = {1,3,1};
    int expected2 = 1;
    vector<int> inputs3 = {100};
    int expected3 = 1;
    Solution sol = Solution();

    printf("output1: %d, expected: %d\n", sol.hIndex(inputs1), expected1);
    printf("output2: %d, expected: %d\n", sol.hIndex(inputs2), expected2);
    printf("output3: %d, expected: %d\n", sol.hIndex(inputs3), expected3);

    return 0;
}
