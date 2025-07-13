package main

import (
	"testing"
)

func TestOddEvenList(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		expected []int
	}{
		{
			name:     "Example 1: [1,2,3,4,5]",
			list:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 3, 5, 2, 4},
		},
		{
			name:     "Example 2: [2,1,3,5,6,4,7]",
			list:     []int{2, 1, 3, 5, 6, 4, 7},
			expected: []int{2, 3, 6, 7, 1, 5, 4},
		},
		{
			name:     "Single element",
			list:     []int{1},
			expected: []int{1},
		},
		{
			name:     "Two elements",
			list:     []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Empty list",
			list:     []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := ListNodeFromIntList(tt.list)
			result := oddEvenList(head)

			// Convert result back to slice for comparison
			var resultSlice []int
			current := result
			for current != nil {
				resultSlice = append(resultSlice, current.Val)
				current = current.Next
			}

			if !result.IsEqual(ListNodeFromIntList(tt.expected)) {
				t.Errorf("oddEvenList() = %v, want %v", resultSlice, tt.expected)
			}
		})
	}
}
