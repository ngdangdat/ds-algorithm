package main

import "testing"

func TestKthSmallestInASortedMatrix(t *testing.T) {
	cases := []struct {
		name     string
		matrix   [][]int
		k        int
		expected int
	}{
		{
			name: "Normal Case 1",
			matrix: [][]int{
				{1, 5, 9},
				{10, 11, 13},
				{12, 13, 15},
			},
			k:        8,
			expected: 13,
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			got := kthSmallest(testCase.matrix, testCase.k)
			if got != testCase.expected {
				t.Errorf("invalid case, expected=%d, got=%d", testCase.expected, got)
			}
		})
	}
}
