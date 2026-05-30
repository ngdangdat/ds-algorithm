package main

import (
	"slices"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	cases := []struct {
		nums     []int
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
	}
	for ci, c := range cases {
		got := productExceptSelf(c.nums)
		if !slices.Equal(got, c.expected) {
			t.Fatalf("Case %d is failed, expected=%v, got=%v", ci, c.expected, got)
		}
	}
}
