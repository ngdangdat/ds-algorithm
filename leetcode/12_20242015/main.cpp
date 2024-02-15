#include<iostream>
#include<vector>
#include<algorithm>

using namespace std;

class Solution {
private:
    int get_larger_candy(int larger_cnt) {
      return (larger_cnt + 1) * larger_cnt / 2;
    }
public:
    int candy(vector<int>& ratings) {
        if (ratings.size() <= 1) return ratings.size();
        int total_candy = ratings.size();
        int larger_count = 0;
        int add_candy = 0;
        int prev = ratings[0];
        for (int i = 1; i < ratings.size(); i++) {
          int curr = ratings[i];
          if (curr > prev) {
            add_candy += 1;
            total_candy += add_candy;
            if (larger_count > 1) {
              total_candy += this->get_larger_candy(larger_count);
              larger_count = 0;
            }
          } else {
              if ((curr < prev) && (add_candy == 0)) {
                larger_count += 1;
              } else {
                larger_count = 0;
              }
              add_candy = 0;
          }
          printf("prev=%d, curr=%d, larger_count=%d, total_candy=%d\n", prev, curr, larger_count, total_candy);
          prev = curr;
        }
        total_candy += this->get_larger_candy(larger_count);
        return total_candy;
    }
};

int main() {
    vector<int> inputs1 = { 1, 6, 10, 8, 7 ,3, 2 };
    int expected1 = 7;

    Solution sol = Solution();

    printf("result1: %d, expected1: %d", sol.candy(inputs1), expected1);
    return 0;
}
