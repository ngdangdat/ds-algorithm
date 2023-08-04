/**
 * @file main.cpp
 * @author your name (you@domain.com)
 * @brief 
 * @version 0.1
 * @date 2023-08-04
 * 
 * @copyright Copyright (c) 2023
 * 
 */
#include<string>

using namespace std;

class Solution {
public:
    bool isSubsequence(string s, string t) {
        if (s.size() == 0) return true;
        if (t.size() == 0 || s.size() > t.size()) return false;
        for (int i = 0, r = 0; i < t.size(); i++) {
            if (t[i] == s[r]) {
                r += 1;
                if (r == s.size()) return true;
            }
        }
        return false;
    }
};

int main() {
    // Input: s = "abc", t = "ahbgdc" true
    // Input: s = "axc", t = "ahbgdc" false
    Solution sol = Solution();
    printf("1: %d", sol.isSubsequence("acb", "ahbgdc"));
    printf("2: %d", sol.isSubsequence("axc", "ahbgdc"));
    return 0;
}
