package main

import "testing"

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1: [1,2,3,4] -> [2,1,4,3]",
			input:    []int{1, 2, 3, 4},
			expected: []int{2, 1, 4, 3},
		},
		{
			name:     "Example 2: empty list -> empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Example 3: [1] -> [1]",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Two nodes: [1,2] -> [2,1]",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "Three nodes: [1,2,3] -> [2,1,3]",
			input:    []int{1, 2, 3},
			expected: []int{2, 1, 3},
		},
		{
			name:     "Five nodes: [1,2,3,4,5] -> [2,1,4,3,5]",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{2, 1, 4, 3, 5},
		},
		{
			name:     "Six nodes: [1,2,3,4,5,6] -> [2,1,4,3,6,5]",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{2, 1, 4, 3, 6, 5},
		},
		{
			name:     "Same values: [1,1,2,2] -> [1,1,2,2]",
			input:    []int{1, 1, 2, 2},
			expected: []int{1, 1, 2, 2},
		},
		{
			name:     "Large values: [100,200,300,400] -> [200,100,400,300]",
			input:    []int{100, 200, 300, 400},
			expected: []int{200, 100, 400, 300},
		},
		{
			name:     "Negative values: [-1,-2,-3,-4] -> [-2,-1,-4,-3]",
			input:    []int{-1, -2, -3, -4},
			expected: []int{-2, -1, -4, -3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create input linked list
			head := ListNodeFromIntList(tt.input)
			
			// Create expected linked list
			expected := ListNodeFromIntList(tt.expected)
			
			// Call the function
			result := swapPairs(head)
			
			// Check if result matches expected
			if !isEqual(result, expected) {
				t.Errorf("swapPairs() = %v, expected %v", result, expected)
			}
		})
	}
}

// Helper function to compare two linked lists
func isEqual(l1, l2 *ListNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	if l1 == nil || l2 == nil {
		return false
	}
	return l1.IsEqual(l2)
}