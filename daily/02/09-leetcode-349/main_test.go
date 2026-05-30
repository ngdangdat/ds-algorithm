package main

import (
	"slices"
	"testing"
)

func TestIntersection(t *testing.T) {
	cases := []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			nums1:    []int{1, 2, 2, 1},
			nums2:    []int{2, 2},
			expected: []int{2},
		},
		{
			nums1:    []int{4, 9, 5},
			nums2:    []int{9, 4, 9, 8, 4},
			expected: []int{4, 9},
		},
	}
	for index, c := range cases {
		got := intersection(c.nums1, c.nums2)
		slices.Sort(got)
		slices.Sort(c.expected)
		if !slices.Equal(c.expected, got) {
			t.Fatalf("Fail case %d expected=%v got=%v", index, c.expected, got)
		}
	}
}
