package main

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "Example 1: 342 + 465 = 807",
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			name:     "Example 2: 0 + 0 = 0",
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			name:     "Example 3: 9999999 + 9999 = 10009998",
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name:     "Different lengths: 99 + 9 = 108",
			l1:       []int{9, 9},
			l2:       []int{9},
			expected: []int{8, 0, 1},
		},
		{
			name:     "Single digits: 5 + 5 = 10",
			l1:       []int{5},
			l2:       []int{5},
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := ListNodeFromIntList(tt.l1)
			l2 := ListNodeFromIntList(tt.l2)
			result := addTwoNumbers(l1, l2)
			
			// Convert result back to slice for comparison
			var resultSlice []int
			current := result
			for current != nil {
				resultSlice = append(resultSlice, current.Val)
				current = current.Next
			}
			
			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("addTwoNumbers() = %v, want %v", resultSlice, tt.expected)
			}
		})
	}
}