package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		expected bool
	}{
		{
			name:     "Example 1: [1,2,2,1] - palindrome",
			list:     []int{1, 2, 2, 1},
			expected: true,
		},
		{
			name:     "Example 2: [1,2] - not palindrome",
			list:     []int{1, 2},
			expected: false,
		},
		{
			name:     "Single element - palindrome",
			list:     []int{1},
			expected: true,
		},
		{
			name:     "Empty list - palindrome",
			list:     []int{},
			expected: true,
		},
		{
			name:     "Odd length palindrome: [1,2,3,2,1]",
			list:     []int{1, 2, 3, 2, 1},
			expected: true,
		},
		{
			name:     "Even length palindrome: [1,2,2,1]",
			list:     []int{1, 2, 2, 1},
			expected: true,
		},
		{
			name:     "Not palindrome: [1,2,3,4,5]",
			list:     []int{1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "All same elements: [1,1,1,1]",
			list:     []int{1, 1, 1, 1},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := ListNodeFromIntList(tt.list)
			result := isPalindrome(head)
			
			if result != tt.expected {
				t.Errorf("isPalindrome() = %v, want %v", result, tt.expected)
			}
		})
	}
}