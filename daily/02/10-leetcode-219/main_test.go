package main

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected bool
	}{
		{
			nums:     []int{1, 2, 3, 1},
			k:        3,
			expected: true,
		},
		{
			nums:     []int{1, 0, 1, 1},
			k:        1,
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 1, 2, 3},
			k:        2,
			expected: false,
		},
	}

	for index, c := range cases {
		got := containsNearbyDuplicate(c.nums, c.k)
		if got != c.expected {
			t.Errorf("Case %d fails, expected=%t got=%t", index, c.expected, got)
		}
	}
}
