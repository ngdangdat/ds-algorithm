#include<iostream>
#include<vector>
#include<map>
#include<tuple>
#include<queue>
#include<algorithm>

using namespace std;

class Solution {
public:
    vector<int> longestObstacleCourseAtEachPosition(vector<int>& obstacles) {
        int n = obstacles.size();
        
        // lis[i] records the lowest increasing sequence of length i + 1.
        vector<int> answer(n, 1), lis;
        
        for (int i = 0; i < n; ++i) {
            // Find the rightmost insertion position idx.
            int idx = upper_bound(lis.begin(), lis.end(), obstacles[i]) - lis.begin();
            if (idx == lis.size())
                lis.push_back(obstacles[i]);
            else
                lis[idx] = obstacles[i];
            answer[i] = idx + 1;
        }
        return answer;
    }
};

int main() {
    // vector<int> inputs = {5,1,5,5,1,3,4,5,1,4};
    vector<int> inputs = {1, 2, 3, 5, 5, 5, 6, 5, 6, 4, 2};
    Solution sol = Solution();
    vector<int> ans = sol.longestObstacleCourseAtEachPosition(inputs);
    for (int i = 0; i < ans.size(); i++) {
        printf("i=%d, inputs[i]=%d, ans[i]=%d\n", i, inputs[i], ans[i]);
    }
    return 0;
}
