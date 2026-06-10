package main

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		wantK    int
		wantHead []int // expected first k elements
	}{
		{"empty", []int{}, 0, []int{}},
		{"single", []int{1}, 1, []int{1}},
		{"all duplicates", []int{2, 2, 2, 2}, 1, []int{2}},
		{"no duplicates", []int{1, 2, 3}, 3, []int{1, 2, 3}},
		{"example1", []int{1, 1, 2}, 2, []int{1, 2}},
		{"example2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{"negatives", []int{-3, -3, -1, 0, 0, 5}, 4, []int{-3, -1, 0, 5}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			k := removeDuplicates(tc.nums)
			if k != tc.wantK {
				t.Fatalf("nums=%v k = %d, want %d", tc.nums, k, tc.wantK)
			}
			for i := range k {
				if tc.nums[i] != tc.wantHead[i] {
					t.Errorf("nums=%v nums[%d] = %d, want %d (got head %v, want %v)",
						tc.nums, i, tc.nums[i], tc.wantHead[i], tc.nums[:k], tc.wantHead)
					break
				}
			}
		})
	}
}
