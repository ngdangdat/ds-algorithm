package main

import (
	"slices"
	"testing"
)

func TestReverseString(t *testing.T) {
	cases := []struct {
		s        []byte
		expected []byte
	}{
		{
			s:        []byte{'h', 'e', 'l', 'l', 'o'},
			expected: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			s:        []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			expected: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}
	for index, c := range cases {
		reverseString(c.s)
		if !slices.Equal(c.expected, c.s) {
			t.Errorf("Case %d failed, expected=%v got=%v", index+1, c.expected, c.s)
		}
	}
}
