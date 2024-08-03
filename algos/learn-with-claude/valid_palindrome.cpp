#include "iostream"
#include <cctype>
#include <string>

using namespace std;

bool is_valid_palindrome(string input) {
  int left = 0, right = input.size() - 1;
  while (left < right) {
    char l = input.at(left), r = input.at(right);
    if (!isalnum(l)) {
      left += 1;
      continue;
    }

    if (!isalnum(r)) {
      right -= 1;
      continue;
    }

    if (tolower(l) != tolower(r)) return false;

    left += 1;
    right -= 1;
  }
  return true;
}

int main() {
  string input = "racecar";
  bool is_valid = is_valid_palindrome(input);
  string is_valid_text = is_valid ? "Yes" : "No";
  cout << "Is valid palindrome?: " << is_valid_text << endl;
  return 0;
}

