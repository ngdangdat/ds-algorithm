package main

import "testing"

func TestMain(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		target   int
		expected bool
	}{
		{
			name:     "Normal Case 1 - target exists in first row",
			input:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   3,
			expected: true,
		},
		{
			name:     "Normal Case 2 - target exists in middle row",
			input:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   11,
			expected: true,
		},
		{
			name:     "Normal Case 3 - target exists in last row",
			input:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   34,
			expected: true,
		},
		{
			name:     "Target not found - too small",
			input:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   0,
			expected: false,
		},
		{
			name:     "Target not found - too large",
			input:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   100,
			expected: false,
		},
		{
			name:     "Target not found - between ranges",
			input:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   8,
			expected: false,
		},
		{
			name:     "Target not found - in valid range but missing",
			input:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   4,
			expected: false,
		},
		{
			name:     "Single row matrix - target found",
			input:    [][]int{{1, 3, 5, 7}},
			target:   5,
			expected: true,
		},
		{
			name:     "Single row matrix - target not found",
			input:    [][]int{{1, 3, 5, 7}},
			target:   4,
			expected: false,
		},
		{
			name:     "Single column matrix - target found",
			input:    [][]int{{1}, {3}, {5}, {7}},
			target:   3,
			expected: true,
		},
		{
			name:     "Single column matrix - target not found",
			input:    [][]int{{1}, {3}, {5}, {7}},
			target:   4,
			expected: false,
		},
		{
			name:     "Single element matrix - target found",
			input:    [][]int{{5}},
			target:   5,
			expected: true,
		},
		{
			name:     "Single element matrix - target not found",
			input:    [][]int{{5}},
			target:   3,
			expected: false,
		},
		{
			name:     "Target at first position",
			input:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   1,
			expected: true,
		},
		{
			name:     "Target at last position",
			input:    [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   60,
			expected: true,
		},
		{
			name:     "Large matrix with target found",
			input:    [][]int{{1, 4, 7, 11}, {2, 5, 8, 12}, {3, 6, 9, 16}, {10, 13, 14, 17}},
			target:   5,
			expected: true,
		},
		{
			name:     "Large matrix with target not found",
			input:    [][]int{{1, 4, 7, 11}, {2, 5, 8, 12}, {3, 6, 9, 16}, {10, 13, 14, 17}},
			target:   15,
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := searchMatrix(testCase.input, testCase.target)
			if got != testCase.expected {
				t.Errorf("invalid, expected=%v, got=%v", testCase.expected, got)
			}
		})
	}
}
