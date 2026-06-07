package main

import (
	"slices"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			nums:     []int{1},
			k:        1,
			expected: []int{1},
		},
		{
			nums:     []int{9, 11},
			k:        2,
			expected: []int{11},
		},
	}
	for index, c := range cases {
		got := maxSlidingWindow(c.nums, c.k)
		if !slices.Equal(got, c.expected) {
			t.Errorf("Case %d fails, expected=%v got=%v", index, c.expected, got)
		}
	}
}
