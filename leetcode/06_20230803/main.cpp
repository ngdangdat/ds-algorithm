/**
 * @file main.cpp
 * @author your name (you@domain.com)
 * @brief https://leetcode.com/problems/insert-delete-getrandom-o1/?envType=study-plan-v2&envId=top-interview-150
 * @version 0.1
 * @date 2023-08-03
 * 
 * @copyright Copyright (c) 2023
 * 
 */
#include<iostream>
#include<vector>
#include<algorithm>
#include<map>
#include<chrono>
#include<cstdlib>
#include <bits/stdc++.h>

using namespace std;

class RandomizedSet {
private:
    unordered_set<int> value_set;
    bool is_existing(int val) {
        unordered_set<int>::iterator find_it = this->value_set.find(val);
        if (find_it == this->value_set.end()) return false;
        return true;
    }

public:
    RandomizedSet() {
        this->value_set = {};
    }
    
    bool insert(int val) {
        if (this->is_existing(val)) {
            return false;
        }
        // this->existing_map.insert(make_pair(val, true));
        this->value_set.insert(val);
        return true;
    }
    
    bool remove(int val) {
        if (!this->is_existing(val)) return false;
        // always exists
        this->value_set.erase(val);
        return true;
    }
    
    int getRandom() {
        return *next(this->value_set.begin(), rand() % this->value_set.size());
    }
};


int main() {
    /**
     ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
     [[], [1], [2], [2], [], [1], [2], []]
     Output
     [null, true, false, true, 2, true, false, 2]
     */
    RandomizedSet s = RandomizedSet();
    printf("1 %d, expected: 1\n", s.insert(1));
    printf("2 %d, expected: 0\n", s.remove(2));
    printf("3 %d, expected: 1\n", s.insert(2));
    printf("4 %d, expected: random\n", s.getRandom());
    printf("5 %d, expected: 1\n", s.remove(1));
    printf("6 %d, expected: 0\n", s.insert(2));
    printf("7 %d, expected: 2 (random only)\n", s.getRandom());
    return 0;
}
