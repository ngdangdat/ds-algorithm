#include<iostream>
#include<vector>

using namespace std;

int binarySearch(vector<int>& els, int l, int r, int num) {
    int index = -1;
    if (l < r) {
        int m = l + (r - l) / 2;
        if (els[m] == num)
            return m;
        if (els[m] > num) {
            return binarySearch(els, l, m - 1, num);
        } else if(els[m] < num) {
            return binarySearch(els, m + 1, r, num);
        }
    }

    if (els[l] < num) {
        return l;
    } else if (l > 0 && els[l - 1] < num) {
        return l - 1;
    }

    return -1;
}

int main() {
    vector<int> v = {1, 2, 3, 4, 6};
    int res = binarySearch(v, 0, v.size() - 1, 5);
    printf("res = %d", res);

    return 0;
}