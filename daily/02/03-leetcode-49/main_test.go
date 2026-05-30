package main

import (
	"cmp"
	"slices"
	"testing"
)

func groupSortFunc(a, b []string) int {
	if len(a) == len(b) {
		return 0
	}
	if len(a) == 0 {
		return -1
	}
	if len(b) == 0 {
		return 1
	}
	return cmp.Compare(a[0], b[0])
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
	}

	for i, c := range cases {
		got := groupAnagrams(c.strs)
		slices.SortFunc(got, groupSortFunc)
		slices.SortFunc(c.expected, groupSortFunc)
		for gi, g := range got {
			e := c.expected[gi]
			if slices.Compare(e, g) > 0 {
				t.Fatalf("Failed case %d, got=%v, expected=%v", i, got, c.expected)
			}
		}
	}

}
