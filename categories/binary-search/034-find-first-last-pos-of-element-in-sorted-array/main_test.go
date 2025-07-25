package main

import (
	"slices"
	"testing"
)

func TestSearchRange(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Normal Case 1",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   8,
			expected: []int{3, 4},
		},
		{
			name:     "Not Found Case 1",
			nums:     []int{5, 7, 7, 8, 8, 10},
			target:   6,
			expected: []int{-1, -1},
		},
		{
			name:     "Not Found Case 2",
			nums:     []int{},
			target:   0,
			expected: []int{-1, -1},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := searchRange(testCase.nums, testCase.target)
			if !slices.Equal(got, testCase.expected) {
				t.Errorf("invalid case, name=%s, expected=%v, got=%v", testCase.name, testCase.expected, got)
			}
		})
	}
}
