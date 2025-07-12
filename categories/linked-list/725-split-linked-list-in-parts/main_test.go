package main

import (
	"testing"
)

func TestSplitListToParts(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		k        int
		expected [][]int
	}{
		{
			name:     "Example 1: [1,2,3] k=5",
			list:     []int{1, 2, 3},
			k:        5,
			expected: [][]int{{1}, {2}, {3}, {}, {}},
		},
		{
			name:     "Example 2: [1,2,3,4,5,6,7,8,9,10] k=3",
			list:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			k:        3,
			expected: [][]int{{1, 2, 3, 4}, {5, 6, 7}, {8, 9, 10}},
		},
		{
			name:     "Example 3: [0,1,2,3,4] k=3",
			list:     []int{0, 1, 2, 3, 4},
			k:        3,
			expected: [][]int{{0, 1}, {2, 3}, {4}},
		},
		{
			name:     "Empty list",
			list:     []int{},
			k:        3,
			expected: [][]int{{}, {}, {}},
		},
		{
			name:     "Single element",
			list:     []int{1},
			k:        1,
			expected: [][]int{{1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := ListNodeFromIntList(tt.list)
			result := splitListToParts(head, tt.k)
			// Convert result to [][]int for comparison
			var resultSlices [][]int
			for _, part := range result {
				var partSlice []int
				current := part
				for current != nil {
					partSlice = append(partSlice, current.Val)
					current = current.Next
				}
				resultSlices = append(resultSlices, partSlice)
			}

			if len(resultSlices) != len(tt.expected) {
				t.Errorf("splitListToParts() returned %d parts, want %d", len(resultSlices), len(tt.expected))
				return
			}

			for i := range resultSlices {
				if len(resultSlices[i]) != len(tt.expected[i]) {
					t.Errorf("Part %d: got length %d, want %d", i, len(resultSlices[i]), len(tt.expected[i]))
					continue
				}
				for j := range resultSlices[i] {
					if resultSlices[i][j] != tt.expected[i][j] {
						t.Errorf("Part %d[%d]: got %d, want %d", i, j, resultSlices[i][j], tt.expected[i][j])
					}
				}
			}
		})
	}
}
