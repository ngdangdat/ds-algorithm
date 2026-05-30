package main

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected []int
	}{
		{
			nums:     []int{1, 1, 1, 2, 2, 3},
			k:        2,
			expected: []int{1, 2},
		},
		{
			nums:     []int{1, 2, 1, 2, 1, 2, 3, 1, 3, 2},
			k:        2,
			expected: []int{1, 2},
		},
		{
			nums:     []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6},
			k:        10,
			expected: []int{1, 2, 5, 3, 6, 7, 4, 8, 10, 11},
		},
	}

	for idx, c := range cases {
		got := topKFrequent(c.nums, c.k)
		slices.Sort(got)
		slices.Sort(c.expected)
		if !slices.Equal(got, c.expected) {
			t.Fatalf("Failed %d case, expected=%v, got=%v", idx, c.expected, got)
		}
	}
}
