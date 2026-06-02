package main

import (
	"slices"
	"testing"
)

func TestFourSum(t *testing.T) {
	cases := []struct {
		nums     []int
		target   int
		expected [][]int
	}{
		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			expected: [][]int{
				{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1},
			},
		},
		{
			nums:     []int{2, 2, 2, 2, 2},
			target:   8,
			expected: [][]int{{2, 2, 2, 2}},
		},
	}

	for index, c := range cases {
		got := fourSum(c.nums, c.target)
		if len(got) != len(c.expected) {
			t.Errorf("case %d, len mismatch, len(got)=%d len(expected)=%d", index, len(got), len(c.expected))
			continue
		}
		// Normalize each triple to ascending order so comparison is order-independent.
		for _, g := range got {
			slices.Sort(g)
		}
		for _, e := range c.expected {
			slices.Sort(e)
		}
		// Sort the outer lists with a full lexicographic key so ties (shared first
		// element) order deterministically in both got and expected.
		slices.SortFunc(got, slices.Compare)
		slices.SortFunc(c.expected, slices.Compare)
		if !slices.EqualFunc(got, c.expected, func(a, b []int) bool {
			return slices.Equal(a, b)
		}) {
			t.Errorf("Failed case %d, expected=%v got=%v", index, c.expected, got)
		}
	}
}
