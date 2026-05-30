package main

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	cases := []struct {
		nums     []int
		k        int
		expected bool
	}{}

	for index, c := range cases {
		got := containsNearbyDuplicate(c.nums, c.k)
		if got != c.expected {
			t.Errorf("Case %d fails, expected=%t got=%t", index, c.expected, got)
		}
	}
}
