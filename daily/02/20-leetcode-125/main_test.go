package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		s        string
		expected bool
	}{
		{
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			s:        "race a car",
			expected: false,
		},
		{
			s:        ".,",
			expected: true,
		},
	}
	for index, c := range cases {
		got := isPalindrome(c.s)
		if got != c.expected {
			t.Errorf("Case %d failed, input=[%s] expected=%t got=%t", index+1, c.s, c.expected, got)
		}
	}
}
