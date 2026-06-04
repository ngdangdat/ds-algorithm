package main

import (
	"testing"
)

func TestSubarraySum(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected int
	}{
		{
			nums:     []int{1, 1, 1},
			k:        2,
			expected: 2,
		},
		{
			nums:     []int{1, 2, 3},
			k:        3,
			expected: 2,
		},
		{
			nums:     []int{-1, 1, 0},
			k:        0,
			expected: 3,
		},
	}
	for index, c := range cases {
		got := subarraySum(c.nums, c.k)
		if c.expected != got {
			t.Errorf("Case %d failed, expected=%d got=%d", index, c.expected, got)
		}
	}
}
