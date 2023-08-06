/**
 * @file main.cpp
 * @author your name (you@domain.com)
 * @brief https://leetcode.com/problems/gas-station/?envType=study-plan-v2&envId=top-interview-150
 * @version 0.1
 * @date 2023-08-06
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
    int canCompleteCircuit(vector<int>& gas, vector<int>& cost) {
        int sum_gas = 0, sum_cost = 0;
        for (int i = 0; i < gas.size(); i++) {
            sum_gas += gas[i];
            sum_cost += cost[i];
        }
        for (int i = 0; i < gas.size(); i++) {
            if (sum_gas - (sum_cost - cost[i])) {

            }
        }
    }
};

int main() {
    Solution sol = Solution();
    vector<int> gas1 = {1,2,3,4,5};
    vector<int> cost1 = {3,4,5,1,2};
    int result1 = sol.canCompleteCircuit(gas1, cost1);
    int expected1 = 3;

    printf("result1=%d, expected1=%d", result1, expected1);

    vector<int> gas2 = {2,3,4};
    vector<int> cost2 = {3,4,3};
    int result2 = sol.canCompleteCircuit(gas2, cost2);
    int expected2 = -1;

    printf("result2=%d, expected2=%d", result2, expected2); 
    return 0;
}
