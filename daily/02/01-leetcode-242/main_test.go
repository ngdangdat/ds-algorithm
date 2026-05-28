package main

import "testing"

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		s        string
		t        string
		expected bool
	}{
		{
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			s:        "rat",
			t:        "car",
			expected: false,
		},
	}

	for i, c := range cases {
		got := isAnagram(c.s, c.t)
		if got != c.expected {
			t.Fatalf("failed case %d, expected=%t got=%t", i+1, c.expected, got)
		}
	}

}
