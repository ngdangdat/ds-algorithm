#include<iostream>
#include<vector>
#include<math.h>

using namespace std;


class Solution {
public:
    int hIndex(vector<int>& citations) {
        int sum = 0;
        for (int i = 0; i < citations.size(); i++) {
            sum += citations[i];
        }
        float avg = (float) sum / citations.size();
        return (int) std::round(avg);
    }
};

int main() {
    vector<int> inputs = {1, 3, 1};
    int expected = 1;

    Solution sol = Solution();

    printf("hindex=[%d], expected=[%d]", sol.hIndex(inputs), expected);
}
