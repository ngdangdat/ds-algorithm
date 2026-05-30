package main

import "testing"

func TestLongestConsecutive(t *testing.T) {
	cases := []struct {
		nums     []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{1, 0, 2, 1}, 3},
	}
	for index, c := range cases {
		got := longestConsecutive(c.nums)
		if got != c.expected {
			t.Fatalf("Failed case %d, expected=%d got=%d", index, c.expected, got)
		}
	}

}
