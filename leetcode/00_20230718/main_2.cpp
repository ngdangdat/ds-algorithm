#include<iostream>
#include<vector>
#include<map>
#include<tuple>
#include<queue>

using namespace std;

class Solution {

public:
    vector<int> longestObstacleCourseAtEachPosition(vector<int>& obstacles) {
        vector<int> ans(obstacles.size(), 0);

        /**
         * {"maxLength": <num, startIndex>}
         * 
         */
        auto compare = [](tuple<int, int, int> l, tuple<int, int, int> r) { return get<0>(l) < get<0>(r); };
        priority_queue<tuple<int, int, int>, vector<tuple<int, int, int>>, decltype(compare)> numMaxLengthPairPq(compare);
        vector<int> sortedProcessedNums;
        for (int i = 0; i < obstacles.size(); i++) {
            int num = obstacles[i];
            if (i == 0) {
                ans[i] = 1;
                numMaxLengthPairPq.push(make_tuple(1, num, 0));
                continue;
            }
            int prev = obstacles[i - 1];
            int _ans = ans[i - 1];
            vector<tuple<int, int, int>> poppedEls;
            bool done = false;
            while(!done && numMaxLengthPairPq.size() > 0) {
                tuple<int, int, int> el = numMaxLengthPairPq.top();
                int currNum = get<1>(el);
                int index = i;
                if (num >= currNum) {
                    int index = get<2>(el);
                    _ans = ans[index] + 1;
                    done = true;
                    break;
                }
                numMaxLengthPairPq.pop();
                poppedEls.push_back(el);
            }
            if (!done) {
                _ans = 1;
            }

            if (poppedEls.size() > 0) {
                for (int j = 0; j < poppedEls.size(); j++) {
                    numMaxLengthPairPq.push(poppedEls[j]);
                }
            }
            ans[i] = _ans;
            numMaxLengthPairPq.push(make_tuple(_ans, num, i));
        }
        return ans;
    }
};

int main() {
    vector<int> inputs = {5,1,5,5,1,3,4,5,1,4};
    Solution sol = Solution();
    vector<int> ans = sol.longestObstacleCourseAtEachPosition(inputs);
    for (int i = 0; i < ans.size(); i++) {
        printf("i=%d, inputs[i]=%d, ans[i]=%d\n", i, inputs[i], ans[i]);
    }
    return 0;
}
