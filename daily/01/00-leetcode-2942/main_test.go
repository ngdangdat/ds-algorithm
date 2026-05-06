package main

import (
	"slices"
	"testing"
)

type testcase struct {
	words    []string
	x        byte
	expected []int
}

func TestFindWordsContaining(t *testing.T) {
	cases := []testcase{
		{
			words:    []string{"leet", "code"},
			x:        byte('e'),
			expected: []int{0, 1},
		},
		{
			words:    []string{"abc", "bcd", "aaaa", "cbc"},
			x:        byte('a'),
			expected: []int{0, 2},
		},
	}

	for i, tCase := range cases {
		res := findWordsContaining(tCase.words, tCase.x)
		if !slices.Equal(res, tCase.expected) {
			t.Errorf("Case %d failed, words=%v x=%v, want=%v got=%v\n", i+1, tCase.words, tCase.x, tCase.expected, res)
		}
	}
}
