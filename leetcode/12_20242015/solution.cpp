#include<iostream>
#include<vector>
#include<algorithm>

using namespace std;

class Solution {
public:
    int candy(vector<int>& ratings) {
      if (ratings.size() == 0) return 0;
      int total = 1;
      int up = 0, down = 0, peak = 0;
      int prev = ratings[0];
      for (int i = 1; i < ratings.size(); i++) {
        int curr = ratings[i];
        if (curr > prev) {
          down = 0;
          up += 1;
          peak = up;
          total += 1 + up;
        } else if (curr == prev) {
          up = 0;
          peak = 0;
          down = 0;
          total += 1;
        } else {
          down += 1;
          up = 0;
          int add = 0;
          if (peak < down) add = 1;
          total += down + add;
        }
        prev = curr;
      }
      return total;
    }
};

int main() {
    vector<int> inputs1 = { 1, 2, 2 };
    int expected1 = 4;

    Solution sol = Solution();

    printf("result1: %d, expected1: %d", sol.candy(inputs1), expected1);
    return 0;
}
