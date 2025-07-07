#include <iostream>
#include <unordered_map>
#include <vector>
using namespace std;

vector<int> twoSum1(vector<int> inputs, int target) {
  vector<int> result;
  for (int i = 0; i < inputs.size(); i++) {
    int ipI = inputs[i];
    if (ipI > target) continue;
    for (int j = 0; j < inputs.size(); j++) {
      if (i == j) continue;
      int ipJ = inputs[j];
      if (ipJ > target) continue;
      if (ipJ + ipI == target) {
        result.push_back(i);
        result.push_back(j);
        break;
      }
    }
    if (result.size() != 0) break;
  }
  return result;
}

vector<int> twoSum2(vector<int> inputs, int target) {
  unordered_map<int, int> numMap = {};
  for (int i = 0; i < inputs.size(); i++) {
    int num = inputs.at(i);
    int complement = target - num;
    if (numMap.find(complement) != numMap.end()) {
      return {numMap.at(complement), i};
    }
    numMap[num] = i;
  }
  return {};
}


int main() {
  vector<int> v = {2, 7, 11, 15};
  int target = 9;
  
  vector<int> result = twoSum2(v, target);
  for (int i = 0; i < result.size(); i++) {
    cout << result[i] << endl;
  }
  return 0;
}
