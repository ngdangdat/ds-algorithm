package main

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		s        string
		expected int
	}{
		{
			s:        "abcabcbb",
			expected: 3,
		},
		{
			s:        "dvdf",
			expected: 3,
		},
	}
	for index, c := range cases {
		fmt.Printf("Case %d %s\n", index+1, c.s)
		got := lengthOfLongestSubstring(c.s)
		if got != c.expected {
			t.Errorf("Case %d failed, expected=%d got=%d", index+1, c.expected, got)
		}
	}
}
