package main

import (
	"reflect"
	"testing"
)

func TestSortList(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		expected []int
	}{
		{
			name:     "Example 1: [4,2,1,3]",
			list:     []int{4, 2, 1, 3},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Example 2: [-1,5,3,4,0]",
			list:     []int{-1, 5, 3, 4, 0},
			expected: []int{-1, 0, 3, 4, 5},
		},
		{
			name:     "Empty list",
			list:     []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			list:     []int{1},
			expected: []int{1},
		},
		{
			name:     "Already sorted",
			list:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted",
			list:     []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Duplicates",
			list:     []int{3, 1, 2, 3, 1},
			expected: []int{1, 1, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := ListNodeFromIntList(tt.list)
			result := sortList(head)
			
			// Convert result back to slice for comparison
			var resultSlice []int
			current := result
			for current != nil {
				resultSlice = append(resultSlice, current.Val)
				current = current.Next
			}
			
			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("sortList() = %v, want %v", resultSlice, tt.expected)
			}
		})
	}
}