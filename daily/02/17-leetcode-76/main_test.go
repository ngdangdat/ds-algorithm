package main

import "testing"

func TestMinWindow(t *testing.T) {
	cases := []struct {
		s        string
		t        string
		expected string
	}{
		{
			s:        "ADOBECODEBANC",
			t:        "ABC",
			expected: "BANC",
		},
	}

	for idx, c := range cases {
		got := minWindow(c.s, c.t)
		if got != c.expected {
			t.Errorf("Failed %d case, expected=%s, got=%s", idx, c.expected, got)
		}
	}
}
