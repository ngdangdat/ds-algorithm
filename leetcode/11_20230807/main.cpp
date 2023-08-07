#include<iostream>
#include<vector>
#include<algorithm>

using namespace std;

class Solution {
public:
    int candy(vector<int>& ratings) {
        int min_index = (int) (min_element(ratings.begin(), ratings.end()) - ratings.begin());
        vector<int> candies(ratings.size(), 1);
        int total_candy = 1;
        for (int i = min_index + 1; i < ratings.size(); i++) {
            if (ratings[i] <= ratings[i - 1] && candies[i - 1] > 1) {
                candies[i] = candies[i - 1] - 1;
            }
            if (ratings[i] > ratings[i - 1]) {
                candies[i] = candies[i - 1] + 1;
            }
            total_candy += candies[i];
        }
        for (int i = min_index - 1; i >= 0; i--) {
            if (ratings[i] <= ratings[i + 1] && candies[i + 1] > 1) {
                candies[i] = candies[i + 1] - 1;
            }
            if (ratings[i] > ratings[i + 1]) {
                candies[i] = candies[i + 1] + 1;
            }
            total_candy += candies[i];
        }
        return total_candy;
    }
};

int main() {
    vector<int> inputs1 = {1, 0, 2};
    int expected1 = 5;

    Solution sol = Solution();

    printf("result1: %d, expected1: %d", sol.candy(inputs1), expected1);
    return 0;
}
