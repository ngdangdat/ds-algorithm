package main

import (
	"slices"
	"strings"
	"testing"
)

func sortOuter(a, b []string) int {
	if len(a) == 0 && len(b) == 0 {
		return 0
	}
	return strings.Compare(a[0], b[0])
}

func TestAnagramGroup(t *testing.T) {
	cases := []struct {
		strs     []string
		expected [][]string
	}{
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			expected: [][]string{
				{"tan", "nat"},
				{"eat", "tea", "ate"},
				{"bat"},
			},
		},
		{
			strs:     []string{""},
			expected: [][]string{{""}},
		},
		{
			strs:     []string{"a"},
			expected: [][]string{{"a"}},
		},
	}

	for i, c := range cases {
		got := groupAnagrams(c.strs)
		for _, g := range got {
			slices.Sort(g)
		}
		for _, e := range c.expected {
			slices.Sort(e)
		}
		slices.SortFunc(got, sortOuter)
		slices.SortFunc(c.expected, sortOuter)
		eq := slices.EqualFunc(got, c.expected, func(a, b []string) bool {
			return slices.Equal(a, b)
		})
		if !eq {
			t.Fatalf("Case %d failed, got=%v, expected=%v\n", i, got, c.expected)
		}
	}

}
